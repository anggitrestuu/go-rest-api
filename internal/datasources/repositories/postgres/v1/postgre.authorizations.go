package v1

import (
	"context"

	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/internal/datasources/records"
	paginate "github.com/anggitrestuu/go-rest-api/pkg/paginate"
	"gorm.io/gorm"
)

type postgresAuthorizationRepository struct {
	conn *gorm.DB
}

func NewAuthorizationRepository(conn *gorm.DB) V1Domains.AuthorizationRepository {
	return &postgresAuthorizationRepository{
		conn: conn,
	}
}

func (r *postgresAuthorizationRepository) Store(ctx context.Context, inDom *V1Domains.AuthorizationDomain) (outDom V1Domains.AuthorizationDomain, err error) {
	authorizationRecord := records.FromAuthorizationV1Domain(inDom)
	result := r.conn.WithContext(ctx).Create(&authorizationRecord)
	if result.Error != nil {
		return V1Domains.AuthorizationDomain{}, result.Error
	}

	return authorizationRecord.ToV1Domain(), nil
}

func (r *postgresAuthorizationRepository) GetByID(ctx context.Context, id int) (outDomain V1Domains.AuthorizationDomain, err error) {
	var authorizationRecord records.Authorizations
	result := r.conn.WithContext(ctx).Where("id = ?", id).First(&authorizationRecord)
	if result.Error != nil {
		return V1Domains.AuthorizationDomain{}, result.Error
	}

	return authorizationRecord.ToV1Domain(), nil
}

func (r *postgresAuthorizationRepository) Update(ctx context.Context, inDom *V1Domains.AuthorizationDomain) (err error) {
	authorizationRecord := records.FromAuthorizationV1Domain(inDom)
	result := r.conn.WithContext(ctx).Save(&authorizationRecord)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *postgresAuthorizationRepository) Delete(ctx context.Context, id int) (err error) {
	result := r.conn.WithContext(ctx).Delete(&records.Authorizations{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *postgresAuthorizationRepository) GetAll(ctx context.Context, params paginate.Params) (outDom paginate.Pagination[V1Domains.AuthorizationDomain], err error) {

	pagination, err := paginate.ToPagination[records.Authorizations](params)

	if err != nil {
		return paginate.Pagination[V1Domains.AuthorizationDomain]{}, err
	}

	// Apply pagination to the database query
	if err := pagination.Paginate(r.conn.WithContext(ctx)); err != nil {
		return paginate.Pagination[V1Domains.AuthorizationDomain]{}, err
	}

	// Transform the pagination items
	result := paginate.TransformPagination(pagination, records.ToAuthorizationV1Domain)

	return result, nil
}
