package responses

import (
	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/pkg/paginate"
)

type RoleResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *RoleResponse) ToV1Domain() V1Domains.RoleDomain {
	return V1Domains.RoleDomain{
		ID:          r.Id,
		Name:        r.Name,
		Description: r.Description,
	}
}

func FromRoleV1Domain(r V1Domains.RoleDomain) RoleResponse {
	return RoleResponse{
		Id:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}
}

type RoleResponses []RoleResponse

func FromRoleV1Domains(r []V1Domains.RoleDomain) RoleResponses {
	var responses RoleResponses
	for _, v := range r {
		responses = append(responses, FromRoleV1Domain(v))
	}
	return responses
}

func PaginationFromRolesV1Domains(u paginate.Pagination[V1Domains.RoleDomain]) paginate.Pagination[RoleResponse] {
	var responses paginate.Pagination[RoleResponse]
	for _, v := range u.Items {
		responses.Items = append(responses.Items, FromRoleV1Domain(v))
	}

	responses.Filters = u.Filters
	responses.Limit = u.Limit
	responses.Page = u.Page
	responses.TotalPages = u.TotalPages
	responses.TotalItems = u.TotalItems

	return responses

}
