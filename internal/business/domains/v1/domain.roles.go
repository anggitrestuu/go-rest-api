package v1

import "context"

type RoleDomain struct {
	ID   int
	Name string
}

type RoleDomainRepository interface {
	GetAll(ctx context.Context) (outDomain []RoleDomain, err error)
}
