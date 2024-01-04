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

func ToArrayOfAuthorizationV1Domain(u *[]Authorizations) []V1Domains.AuthorizationDomain {
	var result []V1Domains.AuthorizationDomain

	for _, val := range *u {
		result = append(result, val.ToV1Domain())
	}

	return result
}

func ToAuthorizationV1Domain(a *Authorizations) V1Domains.AuthorizationDomain {
	// Conversion logic from records.Authorizations to V1Domains.AuthorizationDomain
	return V1Domains.AuthorizationDomain{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
	}
}
