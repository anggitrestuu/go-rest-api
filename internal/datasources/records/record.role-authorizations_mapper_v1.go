package records

import V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"

func (u *RoleAuthorizations) ToV1Domain() V1Domains.RoleAuthorizationsDomain {
	return V1Domains.RoleAuthorizationsDomain{
		RolesID:          u.RolesID,
		AuthorizationsID: u.AuthorizationsID,
	}
}

func FromRoleAuthorizationsV1Domain(u *V1Domains.RoleAuthorizationsDomain) RoleAuthorizations {
	return RoleAuthorizations{
		RolesID:          u.RolesID,
		AuthorizationsID: u.AuthorizationsID,
	}
}
