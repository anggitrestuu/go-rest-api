package v1

import (
	"context"
	"fmt"

	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/internal/datasources/records"
	gorm_paginations "github.com/anggitrestuu/go-rest-api/pkg/paginations"
	"gorm.io/gorm"
)

type postgreUserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) V1Domains.UserRepository {
	return &postgreUserRepository{
		conn: conn,
	}
}

func (r *postgreUserRepository) Store(ctx context.Context, inDom *V1Domains.UserDomain) (err error) {
	userRecord := records.FromUsersV1Domain(inDom)
	// Using GORM's Create method to insert a new record
	result := r.conn.WithContext(ctx).Create(&userRecord)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *postgreUserRepository) GetByEmail(ctx context.Context, inDom *V1Domains.UserDomain) (outDomain V1Domains.UserDomain, err error) {
	userRecord := records.FromUsersV1Domain(inDom)

	paginationParams := gorm_paginations.PaginationParams{Page: 1, Limit: 10, SortBy: "name:desc", Filters: ""}

	fmt.Println("paginationParams", paginationParams)
	paginatedResult, err := gorm_paginations.Paginate(r.conn.WithContext(ctx), paginationParams, userRecord)
	if err != nil {
		fmt.Println("errorPaginate", err)
		return V1Domains.UserDomain{}, err
	}

	fmt.Println("paginatedResult", paginatedResult)

	// Using GORM's First method to fetch the first record matching the email
	result := r.conn.WithContext(ctx).Where("email = ?", userRecord.Email).First(&userRecord)
	if result.Error != nil {
		return V1Domains.UserDomain{}, result.Error
	}

	return userRecord.ToV1Domain(), nil
}

// checkIfEmailExist is a helper function to check if email already exist in database
func (r *postgreUserRepository) CheckIfEmailExist(ctx context.Context, inDom *V1Domains.UserDomain) (outDomain V1Domains.UserDomain, err error) {
	userRecord := records.FromUsersV1Domain(inDom)

	// Using GORM's First method to fetch the first record matching the email
	result := r.conn.WithContext(ctx).Where("email = ?", userRecord.Email).First(&userRecord)
	if result.Error != nil {
		return V1Domains.UserDomain{}, result.Error
	}

	return userRecord.ToV1Domain(), nil
}

func (r *postgreUserRepository) ChangeActiveUser(ctx context.Context, inDom *V1Domains.UserDomain) (err error) {
	userRecord := records.FromUsersV1Domain(inDom)

	// Using GORM's Model and Update method to update the active status
	result := r.conn.WithContext(ctx).Model(&userRecord).Where("id = ?", userRecord.ID).Update("active", userRecord.Active)
	return result.Error
}
