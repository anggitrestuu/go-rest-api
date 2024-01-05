package requests

import V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"

type ProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

// Mapping General Request to Domain Role
func (r ProductRequest) ToV1Domain() *V1Domains.ProductDomain {
	return &V1Domains.ProductDomain{
		Name:        r.Name,
		Description: r.Description,
	}
}
