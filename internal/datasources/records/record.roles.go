package records

import "gorm.io/gorm"

type Roles struct {
	gorm.Model
	ID             int               `gorm:"primaryKey;autoIncrement;"`
	Name           string            `gorm:"type:varchar(200);unique;"`
	Description    string            `gorm:"type:varchar(200);"`
	Accounts       []*Accounts       `gorm:"many2many:account_roles;"`
	Authorizations []*Authorizations `gorm:"many2many:role_authorizations;"`
}
