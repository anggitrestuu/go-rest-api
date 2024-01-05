package v1

import (
	"context"
	"errors"
	"fmt"

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

	// check if role and authorization exists
	var role records.Roles
	var authorization records.Authorizations

	result := r.conn.WithContext(ctx).First(&role, roleID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("role with ID %d not found", roleID)
		}
		return result.Error
	}

	result = r.conn.WithContext(ctx).First(&authorization, authorizationID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("authorization with ID %d not found", authorizationID)
		}
		return result.Error
	}

	result = r.conn.WithContext(ctx).Create(&roleAuthorizationRecord)
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
	var role records.Roles
	var authDomains []V1Domains.AuthorizationDomain

	// Fetch role with related authorizations
	result := r.conn.Preload("Authorizations").First(&role, roleID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("role with ID %d not found", roleID)
		}
		return nil, result.Error
	}

	// Convert to domain model (assuming you have a conversion function)
	for _, auth := range role.Authorizations {
		authDomains = append(authDomains, auth.ToV1Domain())
	}

	return authDomains, nil
}
