package v1

import (
	"context"

	"github.com/anggitrestuu/go-rest-api/pkg/paginate"
)

// ProductDomain represents the domain model for roles.
type ProductDomain struct {
	ID             int
	Name           string
	Description    string
	Authorizations []AuthorizationDomain
}

// ProductUseCase represents the roles useCase contract.
type ProductUseCase interface {
	Store(ctx context.Context, inDom *ProductDomain) (outDom ProductDomain, statusCode int, err error)
	GetByID(ctx context.Context, id int) (outDom ProductDomain, statusCode int, err error)
	Update(ctx context.Context, inDom *ProductDomain) (outDom ProductDomain, statusCode int, err error)
	Delete(ctx context.Context, id int) (statusCode int, err error)
	GetAll(ctx context.Context, params paginate.Params) (outDom paginate.Pagination[ProductDomain], statusCode int, err error)
}

// ProductRepository represents the roles repository contract.
type ProductRepository interface {
	Store(ctx context.Context, inDom *ProductDomain) (outDom ProductDomain, err error)
	GetByID(ctx context.Context, id int) (outDom ProductDomain, err error)
	Update(ctx context.Context, inDom *ProductDomain) (err error)
	Delete(ctx context.Context, id int) (err error)
	GetAll(ctx context.Context, params paginate.Params) (outDom paginate.Pagination[ProductDomain], err error)
}
