package paginations

import (
	"gorm.io/gorm"
)

// PaginationParams represents the parameters needed for paginating queries
type PaginationParams struct {
	Page    int    `json:"page" example:"1"`                              // the page number
	Limit   int    `json:"limit" example:"10"`                            // the number of items per page
	SortBy  string `json:"sort_by" example:"id:desc;name:desc;color:asc"` // optional: ex: "id:desc;name:desc;color:asc"
	Filters string `json:"filters" example:"$eq:account.id=1;"`           // optional
}

// ValidateAndSetDefaults sets default values and validates PaginationParams
func (p *PaginationParams) ValidateAndSetDefaults() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 || p.Limit > 100 {
		p.Limit = 100
	}
}

// PaginatedResult represents the result of a paginated query
type PaginatedResult[T any] struct {
	Items      []T   // the items of the current page
	TotalItems int64 // total number of items
	TotalPages int   // total number of pages
	Page       int   // current page number
	Limit      int   // number of items per page
}

// Paginate performs the paginated query
func Paginate[T any](db *gorm.DB, params PaginationParams, model T) (*PaginatedResult[T], error) {
	var items []T
	result := &PaginatedResult[T]{Page: params.Page, Limit: params.Limit}

	params.ValidateAndSetDefaults()

	// Calculate the total items and pages
	if err := db.Model(&model).Count(&result.TotalItems).Error; err != nil {
		return nil, err
	}
	result.TotalPages = int((result.TotalItems + int64(params.Limit) - 1) / int64(params.Limit))

	// Apply pagination
	offset := (params.Page - 1) * params.Limit
	query := db.Offset(offset).Limit(params.Limit)

	// Apply sorting
	if params.SortBy != "" {
		sortOrder := ""
		sortBy := params.SortBy
		for _, v := range sortBy {
			if v == ';' {
				sortOrder += ", "
			} else if v == ':' {
				sortOrder += " "
			} else {
				sortOrder += string(v)
			}
		}
	}

	parseFilters, err := ParseFilterFromString(params.Filters)
	if err != nil {
		return nil, err
	}

	query = ApplyFilters(db, parseFilters)

	// Execute query
	if err := query.Find(&items).Error; err != nil {
		return nil, err
	}
	result.Items = items

	return result, nil
}

// example using Paginate
