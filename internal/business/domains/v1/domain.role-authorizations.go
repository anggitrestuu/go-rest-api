package v1

import "context"

type RoleAuthorizationsDomain struct {
	RolesID          int
	AuthorizationsID int
}

// RoleAuthorizationUseCase represents the contract for role-authorization business logic
type RoleAuthorizationUseCase interface {
	AssignAuthorizationToRole(ctx context.Context, roleID, authorizationID int) (outDom RoleAuthorizationsDomain, statusCode int, err error)
	RemoveAuthorizationFromRole(ctx context.Context, roleID, authorizationID int) (outDom RoleAuthorizationsDomain, statusCode int, err error)
	GetAuthorizationsByRoleID(ctx context.Context, roleID int) (outDom []AuthorizationDomain, statusCode int, err error)
}

// RoleAuthorizationRepository represents the contract for role-authorization data operations
type RoleAuthorizationRepository interface {
	AssignAuthorizationToRole(ctx context.Context, roleID, authorizationID int) error
	RemoveAuthorizationFromRole(ctx context.Context, roleID, authorizationID int) error
	GetAuthorizationsByRoleID(ctx context.Context, roleID int) ([]AuthorizationDomain, error)
}
