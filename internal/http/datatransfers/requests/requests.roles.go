package requests

import V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"

type RoleRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

// Mapping General Request to Domain Role
func (role RoleRequest) ToV1Domain() *V1Domains.RoleDomain {
	return &V1Domains.RoleDomain{
		Name:        role.Name,
		Description: role.Description,
	}
}
