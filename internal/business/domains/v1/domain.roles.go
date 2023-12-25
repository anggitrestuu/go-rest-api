package v1

// RoleDomain represents the domain model for roles.
type RoleDomain struct {
	ID             int
	Name           string
	Description    string
	Authorizations []AuthorizationDomain
}
