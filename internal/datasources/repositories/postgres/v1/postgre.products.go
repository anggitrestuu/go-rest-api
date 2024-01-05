package v1

import (
	"context"

	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/internal/datasources/records"
	paginate "github.com/anggitrestuu/go-rest-api/pkg/paginate"
	"gorm.io/gorm"
)

type postgresProductRepository struct {
	conn *gorm.DB
}

func NewProductRepository(conn *gorm.DB) V1Domains.ProductRepository {
	return &postgresProductRepository{
		conn: conn,
	}
}

func (r *postgresProductRepository) Store(ctx context.Context, inDom *V1Domains.ProductDomain) (outDom V1Domains.ProductDomain, err error) {
	authorizationRecord := records.FromProductV1Domain(inDom)
	result := r.conn.WithContext(ctx).Create(&authorizationRecord)
	if result.Error != nil {
		return V1Domains.ProductDomain{}, result.Error
	}

	return authorizationRecord.ToV1Domain(), nil
}

func (r *postgresProductRepository) GetByID(ctx context.Context, id int) (outDomain V1Domains.ProductDomain, err error) {
	var authorizationRecord records.Products
	result := r.conn.WithContext(ctx).Where("id = ?", id).First(&authorizationRecord)
	if result.Error != nil {
		return V1Domains.ProductDomain{}, result.Error
	}

	return authorizationRecord.ToV1Domain(), nil
}

func (r *postgresProductRepository) Update(ctx context.Context, inDom *V1Domains.ProductDomain) (err error) {
	authorizationRecord := records.FromProductV1Domain(inDom)
	result := r.conn.WithContext(ctx).Save(&authorizationRecord)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *postgresProductRepository) Delete(ctx context.Context, id int) (err error) {
	result := r.conn.WithContext(ctx).Delete(&records.Products{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *postgresProductRepository) GetAll(ctx context.Context, params paginate.Params) (outDom paginate.Pagination[V1Domains.ProductDomain], err error) {

	pagination, err := paginate.ToPagination[records.Products](params)

	if err != nil {
		return paginate.Pagination[V1Domains.ProductDomain]{}, err
	}

	// Apply pagination to the database query
	if err := pagination.Paginate(r.conn.WithContext(ctx)); err != nil {
		return paginate.Pagination[V1Domains.ProductDomain]{}, err
	}

	// Transform the pagination items
	result := paginate.TransformPagination(pagination, records.ToProductV1Domain)

	return result, nil
}
