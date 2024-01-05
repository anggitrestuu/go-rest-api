package responses

import (
	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
)

type RoleResponse struct {
	Id             int                     `json:"id"`
	Name           string                  `json:"name"`
	Description    string                  `json:"description"`
	Authorizations []AuthorizationResponse `json:"authorizations"`
}

func (r *RoleResponse) ToV1Domain() V1Domains.RoleDomain {

	var roleAuthorizations []V1Domains.AuthorizationDomain
	for _, val := range r.Authorizations {
		roleAuthorizations = append(roleAuthorizations, val.ToV1Domain())
	}

	return V1Domains.RoleDomain{
		ID:             r.Id,
		Name:           r.Name,
		Description:    r.Description,
		Authorizations: roleAuthorizations,
	}
}

func FromRoleV1Domain(r V1Domains.RoleDomain) RoleResponse {

	var roleAuthorizations []AuthorizationResponse
	for _, val := range r.Authorizations {
		roleAuthorizations = append(roleAuthorizations, FromAuthorizationV1Domain(val))
	}

	return RoleResponse{
		Id:             r.ID,
		Name:           r.Name,
		Description:    r.Description,
		Authorizations: roleAuthorizations,
	}
}
