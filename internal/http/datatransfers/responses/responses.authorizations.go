package responses

import (
	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
)

type AuthorizationResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (u *AuthorizationResponse) ToV1Domain() V1Domains.AuthorizationDomain {
	return V1Domains.AuthorizationDomain{
		ID:          uint(u.ID),
		Name:        u.Name,
		Description: u.Description,
	}
}

func FromAuthorizationV1Domain(u V1Domains.AuthorizationDomain) AuthorizationResponse {
	return AuthorizationResponse{
		ID:          int(u.ID),
		Name:        u.Name,
		Description: u.Description,
	}
}
