package responses

import (
	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/pkg/paginate"
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

type AuthorizationResponses []AuthorizationResponse

func FromAuthorizationV1Domains(u []V1Domains.AuthorizationDomain) AuthorizationResponses {
	var responses AuthorizationResponses
	for _, v := range u {
		responses = append(responses, FromAuthorizationV1Domain(v))
	}
	return responses
}

func PaginationFromAuthorizationV1Domains(u paginate.Pagination[V1Domains.AuthorizationDomain]) paginate.Pagination[AuthorizationResponse] {
	var responses paginate.Pagination[AuthorizationResponse]
	for _, v := range u.Items {
		responses.Items = append(responses.Items, FromAuthorizationV1Domain(v))
	}

	responses.Filters = u.Filters
	responses.Limit = u.Limit
	responses.Page = u.Page
	responses.TotalPages = u.TotalPages
	responses.TotalItems = u.TotalItems

	return responses

}
