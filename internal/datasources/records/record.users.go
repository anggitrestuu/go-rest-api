package records

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	Username string `json:"username" gorm:"type:varchar(200);"`
	Email    string `json:"email" gorm:"type:varchar(200);"`
	Password string `json:"password" gorm:"type:varchar(200);"`
	Active   bool   `json:"active" gorm:"type:boolean;"`
	RoleId   int    `json:"role_id" gorm:"type:integer;"`
}
