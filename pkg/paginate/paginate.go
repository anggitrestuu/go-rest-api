package paginate

import (
	"github.com/anggitrestuu/go-rest-api/internal/utils"
	"gorm.io/gorm"
)

type Pagination struct {
	Items      interface{} `json:"items"`
	TotalItems int64       `json:"total_items"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	SortBy     string      `json:"sort_by"`
}

type Params struct {
	Page    string `json:"page" example:"1"`                              // the page number
	Limit   string `json:"limit" example:"10"`                            // the number of items per page
	SortBy  string `json:"sort_by" example:"id:desc;name:desc;color:asc"` // optional: ex: "id:desc;name:desc;color:asc"
	Filters string `json:"filters" example:"$eq:account.id=1;"`           // optional
}

type paginate struct {
	limit   int
	page    int
	sortBy  string
	filters string
	model   interface{}
}

type response struct {
	Items      interface{} `json:"items"`
	TotalItems int64       `json:"total_items"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
}

func NewPaginate(params Params, model interface{}) *paginate {
	return &paginate{
		limit:   utils.StringToInt(params.Limit),
		page:    utils.StringToInt(params.Page),
		sortBy:  params.SortBy,
		filters: params.Filters,
		model:   model,
	}

}

func (p *paginate) Paginate(db *gorm.DB) (*gorm.DB, error) {

	result := &response{Page: p.page, Limit: p.limit}

	// Calculate the total items and pages
	if err := db.Model(p.model).Count(&result.TotalItems).Error; err != nil {
		return nil, err
	}

	result.TotalPages = int((result.TotalItems + int64(p.limit) - 1) / int64(p.limit))

	// apply pagination
	if p.page < 1 {
		p.page = 1
	}
	if p.limit < 1 || p.limit > 100 {
		p.limit = 100
	}
	offset := (p.page - 1) * p.limit

	db = db.Offset(offset).Limit(p.limit)

	// Apply sorting
	if p.sortBy != "" {
		sortOrder := ""
		sortBy := p.sortBy
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
	parseFilters, err := ParseFilterFromString(p.filters)

	if err != nil {
		return nil, err
	}

	db = ApplyFilters(db, parseFilters)

	return db, nil

	//if err != nil {
	//	return nil, err
	//}
	//db = ApplyFilters(db, parseFilters)
	//
	//if err = db.Find(p.model).Error; err != nil {
	//	return nil, err
	//}
	//
	//result.Items = p.model
	//
	//return result, nil
}
