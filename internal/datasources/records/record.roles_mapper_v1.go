package records

import V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"

func (u *Roles) ToV1Domain() V1Domains.RoleDomain {
	return V1Domains.RoleDomain{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
	}
}

func FromRoleV1Domain(u *V1Domains.RoleDomain) Roles {
	return Roles{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
	}
}

func ToArrayOfRoleV1Domain(u *[]Roles) []V1Domains.RoleDomain {
	var result []V1Domains.RoleDomain

	for _, val := range *u {
		result = append(result, val.ToV1Domain())
	}

	return result
}
