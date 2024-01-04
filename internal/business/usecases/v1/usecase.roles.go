package v1

import (
	"context"
	"net/http"

	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/pkg/paginate"
)

type roleUseCase struct {
	repo V1Domains.RoleRepository
}

func NewRoleUseCase(repo V1Domains.RoleRepository) V1Domains.RoleUseCase {
	return &roleUseCase{
		repo: repo,
	}
}

func (roleUC *roleUseCase) Store(ctx context.Context, inDom *V1Domains.RoleDomain) (outDom V1Domains.RoleDomain, statusCode int, err error) {
	outDom, err = roleUC.repo.Store(ctx, inDom)
	if err != nil {
		return V1Domains.RoleDomain{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusCreated, nil
}

func (roleUC *roleUseCase) GetByID(ctx context.Context, id int) (outDom V1Domains.RoleDomain, statusCode int, err error) {
	outDom, err = roleUC.repo.GetByID(ctx, id)

	if err != nil {
		return V1Domains.RoleDomain{}, http.StatusNotFound, err
	}

	return outDom, http.StatusOK, nil
}

func (roleUC *roleUseCase) Update(ctx context.Context, inDom *V1Domains.RoleDomain) (outDom V1Domains.RoleDomain, statusCode int, err error) {
	err = roleUC.repo.Update(ctx, inDom)
	if err != nil {
		return V1Domains.RoleDomain{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}

func (roleUC *roleUseCase) Delete(ctx context.Context, id int) (statusCode int, err error) {
	err = roleUC.repo.Delete(ctx, id)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}

func (roleUC *roleUseCase) GetAll(ctx context.Context, params paginate.Params) (outDom paginate.Pagination[V1Domains.RoleDomain], statusCode int, err error) {
	outDom, err = roleUC.repo.GetAll(ctx, params)
	if err != nil {
		return paginate.Pagination[V1Domains.RoleDomain]{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}
