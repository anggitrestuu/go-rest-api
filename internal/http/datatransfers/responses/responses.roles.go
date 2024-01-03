package responses

import (
	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
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
