package records

type Roles struct {
	ID             uint             `gorm:"primaryKey;autoIncrement;"`
	Name           string           `gorm:"type:varchar(200);unique;"`
	Description    string           `gorm:"type:varchar(200);"`
	Accounts       []Accounts       `gorm:"many2many:account_roles;"`
	Authorizations []Authorizations `gorm:"many2many:role_authorizations;"`
}
