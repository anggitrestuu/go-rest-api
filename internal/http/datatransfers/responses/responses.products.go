package responses

import (
	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
)

type ProductResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *ProductResponse) ToV1Domain() V1Domains.ProductDomain {
	return V1Domains.ProductDomain{
		ID:          r.Id,
		Name:        r.Name,
		Description: r.Description,
	}
}

func FromProductV1Domain(r V1Domains.ProductDomain) ProductResponse {
	return ProductResponse{
		Id:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}
}
