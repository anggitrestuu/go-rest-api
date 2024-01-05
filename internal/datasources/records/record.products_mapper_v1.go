package records

import V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"

func (u *Products) ToV1Domain() V1Domains.ProductDomain {
	return V1Domains.ProductDomain{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
	}
}

func FromProductV1Domain(u *V1Domains.ProductDomain) Products {
	return Products{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
	}
}

func ToProductV1Domain(a *Products) V1Domains.ProductDomain {
	return V1Domains.ProductDomain{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
	}
}
