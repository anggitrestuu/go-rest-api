package records

type RoleAuthorizations struct {
	RolesID          int `gorm:"primaryKey;autoIncrement:false;"`
	AuthorizationsID int `gorm:"primaryKey;autoIncrement:false;"`
}
