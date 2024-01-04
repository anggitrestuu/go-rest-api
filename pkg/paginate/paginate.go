package paginate

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// Transformer is a generic function type for transforming domain models to response models.
type Transformer[T any, R any] func(T) R

func TransformPagination[T any, R any](pagination *Pagination[T], transformer func(*T) R) Pagination[R] {
	var newItems []R
	for _, item := range pagination.Items {
		newItems = append(newItems, transformer(&item))
	}

	return Pagination[R]{
		Items:      newItems,
		TotalItems: pagination.TotalItems,
		TotalPages: pagination.TotalPages,
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		SortBy:     pagination.SortBy,
		Filters:    pagination.Filters,
	}
}

type Pagination[T any] struct {
	Items      []T    `json:"items"`
	TotalItems int64  `json:"total_items"`
	TotalPages int    `json:"total_pages"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	SortBy     string `json:"sort_by"`
	Filters    string `json:"filters"`
}

type Params struct {
	Limit   string `form:"limit"`
	Page    string `form:"page"`
	SortBy  string `form:"sort_by"`
	Filters string `form:"filters"`
}

func ToPagination[T any](p Params) (*Pagination[T], error) {
	page, err := strconv.Atoi(p.Page)
	if err != nil {
		page = 1 // default to first page
	}
	limit, err := strconv.Atoi(p.Limit)
	if err != nil {
		limit = 10 // default to 10 items
	}

	return &Pagination[T]{
		Page:    page,
		Limit:   limit,
		SortBy:  p.SortBy,
		Filters: p.Filters,
	}, nil
}

func (p *Pagination[T]) Paginate(db *gorm.DB) error {
	result := db.Model(&[]T{}).Count(&p.TotalItems)
	if result.Error != nil {
		return result.Error
	}

	p.TotalPages = int((p.TotalItems + int64(p.Limit) - 1) / int64(p.Limit))

	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = 10 // default limit if not specified
	}

	offset := (p.Page - 1) * p.Limit
	result = db.Offset(offset).Limit(p.Limit)

	// Apply sorting
	if p.SortBy != "" {

		sortParts := strings.Split(p.SortBy, ";")
		for _, part := range sortParts {
			if part == "" {
				continue
			}

			// Split into operator and value
			operatorValue := strings.SplitN(part, ",", 2)
			if len(operatorValue) != 2 {
				return result.Error
			}

			operator := operatorValue[0]
			value := operatorValue[1]

			result = result.Order(operator + " " + value)
		}

	}

	// apply filters
	parseFilters, err := ParseFilterFromString(p.Filters)
	if err != nil {
		return err
	}

	if len(parseFilters) > 0 {
		result = ApplyFilters(result, parseFilters)
	}

	return result.Find(&p.Items).Error
}
