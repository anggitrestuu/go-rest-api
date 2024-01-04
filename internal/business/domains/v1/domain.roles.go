package v1

import (
	"context"

	"github.com/anggitrestuu/go-rest-api/pkg/paginate"
)

// RoleDomain represents the domain model for roles.
type RoleDomain struct {
	ID             int
	Name           string
	Description    string
	Authorizations []AuthorizationDomain
}

// RoleUseCase represents the roles useCase contract.
type RoleUseCase interface {
	Store(ctx context.Context, inDom *RoleDomain) (outDom RoleDomain, statusCode int, err error)
	GetByID(ctx context.Context, id int) (outDom RoleDomain, statusCode int, err error)
	Update(ctx context.Context, inDom *RoleDomain) (outDom RoleDomain, statusCode int, err error)
	Delete(ctx context.Context, id int) (statusCode int, err error)
	GetAll(ctx context.Context, params paginate.Params) (outDom paginate.Pagination[RoleDomain], statusCode int, err error)
}

// RoleRepository represents the roles repository contract.
type RoleRepository interface {
	Store(ctx context.Context, inDom *RoleDomain) (outDom RoleDomain, err error)
	GetByID(ctx context.Context, id int) (outDom RoleDomain, err error)
	Update(ctx context.Context, inDom *RoleDomain) (err error)
	Delete(ctx context.Context, id int) (err error)
	GetAll(ctx context.Context, params paginate.Params) (outDom paginate.Pagination[RoleDomain], err error)
}
