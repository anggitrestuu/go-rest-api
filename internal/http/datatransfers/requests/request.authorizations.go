package requests

import V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"

// General Request
type AuthorizationRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// Mapping General Request to Domain Authorization
func (auth AuthorizationRequest) ToV1Domain() *V1Domains.AuthorizationDomain {
	return &V1Domains.AuthorizationDomain{
		Name:        auth.Name,
		Description: auth.Description,
	}
}
