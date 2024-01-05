package records

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ID          int    `gorm:"primaryKey;autoIncrement;"`
	Name        string `gorm:"type:varchar(200);"`
	Description string `gorm:"type:varchar(200);"`
}
