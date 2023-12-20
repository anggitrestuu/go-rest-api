package records

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        string         `gorm:"primaryKey;column:id"`
	Username  string         `gorm:"column:username"`
	Email     string         `gorm:"column:email;unique"`
	Password  string         `gorm:"column:password"`
	Active    bool           `gorm:"column:active"`
	RoleId    int            `gorm:"column:role_id"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
