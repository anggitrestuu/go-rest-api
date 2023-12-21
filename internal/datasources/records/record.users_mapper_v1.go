package records

import (
	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
)

func (u *Users) ToV1Domain() V1Domains.UserDomain {
	return V1Domains.UserDomain{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Active:   u.Active,
		RoleID:   u.RoleId,
	}
}

func FromUsersV1Domain(u *V1Domains.UserDomain) Users {
	return Users{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Active:   u.Active,
		RoleId:   u.RoleID,
	}
}

func ToArrayOfUsersV1Domain(u *[]Users) []V1Domains.UserDomain {
	var result []V1Domains.UserDomain

	for _, val := range *u {
		result = append(result, val.ToV1Domain())
	}

	return result
}
