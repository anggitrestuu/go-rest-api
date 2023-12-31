package v1

import (
	"context"

	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/internal/datasources/records"
	paginate "github.com/anggitrestuu/go-rest-api/pkg/paginate"
	"gorm.io/gorm"
)

type postgreRoleRepository struct {
	conn *gorm.DB
}

func NewRoleRepository(conn *gorm.DB) V1Domains.RoleRepository {
	return &postgreRoleRepository{
		conn: conn,
	}
}

func (r *postgreRoleRepository) Store(ctx context.Context, inDom *V1Domains.RoleDomain) (outDom V1Domains.RoleDomain, err error) {
	roleRecord := records.FromRoleV1Domain(inDom)
	// Using GORM's Create method to insert a new record
	result := r.conn.WithContext(ctx).Create(&roleRecord)
	if result.Error != nil {
		return V1Domains.RoleDomain{}, result.Error
	}

	return roleRecord.ToV1Domain(), nil
}

func (r *postgreRoleRepository) GetByID(ctx context.Context, id int) (outDom V1Domains.RoleDomain, err error) {
	roleRecord := records.Roles{}
	result := r.conn.WithContext(ctx).First(&roleRecord, id)

	if result.Error != nil {
		return V1Domains.RoleDomain{}, result.Error
	}

	return roleRecord.ToV1Domain(), nil
}

func (r *postgreRoleRepository) Update(ctx context.Context, inDom *V1Domains.RoleDomain) (err error) {
	roleRecord := records.FromRoleV1Domain(inDom)

	result := r.conn.WithContext(ctx).Save(&roleRecord)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *postgreRoleRepository) Delete(ctx context.Context, id int) (err error) {
	roleRecord := records.Roles{}
	result := r.conn.WithContext(ctx).Delete(&roleRecord, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *postgreRoleRepository) GetAll(ctx context.Context, params paginate.Params) (outDom paginate.Pagination[V1Domains.RoleDomain], err error) {

	pagination, err := paginate.ToPagination[records.Roles](params)
	pagination.Preload("Authorizations")

	if err != nil {
		return paginate.Pagination[V1Domains.RoleDomain]{}, err
	}

	// Apply pagination to the database query
	if err := pagination.Paginate(r.conn.WithContext(ctx)); err != nil {
		return paginate.Pagination[V1Domains.RoleDomain]{}, err
	}

	result := paginate.TransformPagination(
		pagination,
		records.ToRoleV1Domain,
	)

	return result, nil
}
