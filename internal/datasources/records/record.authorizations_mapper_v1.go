package records

import V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"

func (u *Authorizations) ToV1Domain() V1Domains.AuthorizationDomain {
	return V1Domains.AuthorizationDomain{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
	}
}

func FromAuthorizationV1Domain(u *V1Domains.AuthorizationDomain) Authorizations {
	return Authorizations{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
	}
}

func ToAuthorizationV1Domain(a *Authorizations) V1Domains.AuthorizationDomain {
	return V1Domains.AuthorizationDomain{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
	}
}
