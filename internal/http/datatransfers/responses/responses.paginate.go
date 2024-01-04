package responses

import "github.com/anggitrestuu/go-rest-api/pkg/paginate"

// Transformer is a generic function type for transforming domain models to response models.
type Transformer[T any, R any] func(T) R

// TransformPagination takes a paginated result of type T and transforms it to type R.
func TransformPagination[T any, R any](pagination paginate.Pagination[T], transform Transformer[T, R]) paginate.Pagination[R] {
	var transformedItems []R
	for _, item := range pagination.Items {
		transformedItems = append(transformedItems, transform(item))
	}

	return paginate.Pagination[R]{
		Items:      transformedItems,
		TotalItems: pagination.TotalItems,
		TotalPages: pagination.TotalPages,
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		SortBy:     pagination.SortBy,
		Filters:    pagination.Filters,
	}
}
