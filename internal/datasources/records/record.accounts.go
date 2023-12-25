package records

import "gorm.io/gorm"

type Accounts struct {
	gorm.Model
	Username string  `gorm:"type:varchar(200);unique;"` // unique
	Password string  `gorm:"type:varchar(200);"`
	Email    string  `gorm:"type:varchar(200);unique;"` // unique
	Roles    []Roles `gorm:"many2many:account_roles;"`
}
