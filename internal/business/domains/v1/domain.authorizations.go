package v1

import (
	"context"
)

// AuthorizationDomain represents the domain record for authorizations.
type AuthorizationDomain struct {
	ID          uint
	Name        string
	Description string
}

// AuthorizationUseCase represents the authorizations useCase contract.
type AuthorizationUseCase interface {
	Store(ctx context.Context, inDom *AuthorizationDomain) (outDom AuthorizationDomain, statusCode int, err error)
	GetByID(ctx context.Context, id int) (outDom AuthorizationDomain, statusCode int, err error)
	Update(ctx context.Context, inDom *AuthorizationDomain) (outDom AuthorizationDomain, statusCode int, err error)
	Delete(ctx context.Context, id int) (statusCode int, err error)
	GetAll(ctx context.Context) (result any, statusCode int, err error)
}

// AuthorizationRepository represents the authorizations repository contract.
type AuthorizationRepository interface {
	Store(ctx context.Context, inDom *AuthorizationDomain) (outDom AuthorizationDomain, err error)
	GetByID(ctx context.Context, id int) (outDom AuthorizationDomain, err error)
	Update(ctx context.Context, inDom *AuthorizationDomain) (err error)
	Delete(ctx context.Context, id int) (err error)
	GetAll(ctx context.Context, params any) (result any, err error)
}
