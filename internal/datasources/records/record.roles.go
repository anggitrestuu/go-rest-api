package records

import "gorm.io/gorm"

func init() {

}

type Roles struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement;"`
	Name string `json:"name" gorm:"type:varchar(200);"`
}
