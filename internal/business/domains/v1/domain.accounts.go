package v1

// AccountDomain represents the domain record for accounts.
type AccountDomain struct {
	ID          uint
	Name        string
	Description string
	Roles       []RoleDomain
}
