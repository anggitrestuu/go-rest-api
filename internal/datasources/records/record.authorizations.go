package records

import "gorm.io/gorm"

type Authorizations struct {
	gorm.Model
	ID          uint     `gorm:"primaryKey;autoIncrement;"`
	Name        string   `gorm:"type:varchar(200);"`
	Description string   `gorm:"type:varchar(200);"`
	Roles       []*Roles `gorm:"many2many:role_authorizations;"`
}
