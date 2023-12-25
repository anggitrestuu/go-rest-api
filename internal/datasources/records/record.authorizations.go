package records

type Authorizations struct {
	ID          uint    `gorm:"primaryKey;autoIncrement;"`
	Name        string  `gorm:"type:varchar(200);unique;"`
	Description string  `gorm:"type:varchar(200);"`
	Roles       []Roles `gorm:"many2many:role_authorizations;"`
}
