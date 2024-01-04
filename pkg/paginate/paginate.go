package paginate

import (
	"strconv"

	"gorm.io/gorm"
)

type Pagination struct {
	Items      interface{} `json:"items"`
	TotalItems int64       `json:"total_items"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	SortBy     string      `json:"sort_by"`
	Filters    string      `json:"filters"`
}

type Params struct {
	Limit   string `json:"limit"`
	Page    string `json:"page"`
	SortBy  string `json:"sort_by"`
	Filters string `json:"filters"`
}

// params to pagination
func (p *Params) ToPagination(model interface{}) *Pagination {
	page, _ := strconv.Atoi(p.Page)   // Handle error appropriately
	limit, _ := strconv.Atoi(p.Limit) // Handle error appropriately

	return &Pagination{
		Items:      model,
		TotalItems: 0,
		TotalPages: 0,
		Page:       page,
		Limit:      limit,
		SortBy:     p.SortBy,
		Filters:    p.Filters,
	}
}

func (p *Pagination) Paginate(db *gorm.DB) (*gorm.DB, error) {
	// calculate total items

	if err := db.Model(p.Items).Count(&p.TotalItems).Error; err != nil {
		return nil, err
	}

	// calculate total pages
	p.TotalPages = int((p.TotalItems + int64(p.Limit) - 1) / int64(p.Limit))

	// apply pagination
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 || p.Limit > 100 {
		p.Limit = 100
	}

	offset := (p.Page - 1) * p.Limit

	db = db.Offset(offset).Limit(p.Limit)

	// Apply sorting
	if p.SortBy != "" {
		sortOrder := ""
		sortBy := p.SortBy
		for _, v := range sortBy {
			if v == ';' {
				sortOrder += ", "
			} else if v == ':' {
				sortOrder += " "
			} else {
				sortOrder += string(v)
			}
		}
		//query = query.Order(sortOrder)
		db = db.Order(sortOrder)
	}

	// apply filters
	parseFilters, err := ParseFilterFromString(p.Filters)
	if err != nil {
		return nil, err
	}

	if len(parseFilters) > 0 {
		db = ApplyFilters(db, parseFilters)
	}

	return db, nil

}
