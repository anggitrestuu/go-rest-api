package v1

import (
	"context"
	"fmt"
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

func (r *postgresAuthorizationRepository) GetAll(ctx context.Context, params any) (model any, err error) {

	fmt.Println("params", params)

	newParams := paginate.Params{
		Limit:   "10",
		Page:    "1",
		SortBy:  "",
		Filters: "",
	}

	var authorizationRecords []records.Authorizations
	result, err := paginate.NewPaginate(newParams, &authorizationRecords).Paginate(r.conn.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	//result.Items = records.ToArrayOfAuthorizationV1Domain(&authorizationRecords)

	//response.Items = records.ToArrayOfAuthorizationV1Domain(&authorizationRecords)

	return result, nil

	//return records.ToArrayOfAuthorizationV1Domain(&authorizationRecords), nil
}
