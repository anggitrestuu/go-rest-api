package v1

import (
	"context"

	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/internal/datasources/records"

	"gorm.io/gorm"
)

type postgresRoleAuthorizationRepository struct {
	conn *gorm.DB
}

func NewRoleAuthorizationRepository(conn *gorm.DB) V1Domains.RoleAuthorizationRepository {
	return &postgresRoleAuthorizationRepository{
		conn: conn,
	}
}

func (r *postgresRoleAuthorizationRepository) AssignAuthorizationToRole(ctx context.Context, roleID, authorizationID int) error {
	roleAuthorizationRecord := records.RoleAuthorizations{
		RolesID:          roleID,
		AuthorizationsID: authorizationID,
	}

	result := r.conn.WithContext(ctx).Create(&roleAuthorizationRecord)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *postgresRoleAuthorizationRepository) RemoveAuthorizationFromRole(ctx context.Context, roleID, authorizationID int) error {
	result := r.conn.WithContext(ctx).Where("roles_id = ? AND authorizations_id = ?", roleID, authorizationID).Delete(&records.RoleAuthorizations{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *postgresRoleAuthorizationRepository) GetAuthorizationsByRoleID(ctx context.Context, roleID int) ([]V1Domains.AuthorizationDomain, error) {
	var authorizationRecords []records.Authorizations
	result := r.conn.WithContext(ctx).Where("roles_id = ?", roleID).Find(&authorizationRecords)
	if result.Error != nil {
		return nil, result.Error
	}

	var authorizationDomains []V1Domains.AuthorizationDomain
	for _, authorizationRecord := range authorizationRecords {
		authorizationDomains = append(authorizationDomains, authorizationRecord.ToV1Domain())
	}

	return authorizationDomains, nil
}
