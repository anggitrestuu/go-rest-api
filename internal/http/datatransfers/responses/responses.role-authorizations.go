package responses

import (
	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
)

type RoleAuthorizationResponse struct {
	RolesID          int `json:"roles_id"`
	AuthorizationsID int `json:"authorizations_id"`
}

func FromRoleAuthorizationV1Domain(r V1Domains.RoleAuthorizationsDomain) RoleAuthorizationResponse {
	return RoleAuthorizationResponse{
		RolesID:          r.RolesID,
		AuthorizationsID: r.AuthorizationsID,
	}
}
