package v1

import (
	"context"
	"net/http"

	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/pkg/paginate"
)

type authorizationUseCase struct {
	repo V1Domains.AuthorizationRepository
}

func NewAuthorizationUseCase(repo V1Domains.AuthorizationRepository) V1Domains.AuthorizationUseCase {
	return &authorizationUseCase{
		repo: repo,
	}
}

func (authUC *authorizationUseCase) Store(ctx context.Context, inDom *V1Domains.AuthorizationDomain) (outDom V1Domains.AuthorizationDomain, statusCode int, err error) {
	outDom, err = authUC.repo.Store(ctx, inDom)
	if err != nil {
		return V1Domains.AuthorizationDomain{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusCreated, nil
}

func (authUC *authorizationUseCase) GetByID(ctx context.Context, id int) (outDom V1Domains.AuthorizationDomain, statusCode int, err error) {
	outDom, err = authUC.repo.GetByID(ctx, id)
	if err != nil {
		return V1Domains.AuthorizationDomain{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}

func (authUC *authorizationUseCase) Update(ctx context.Context, inDom *V1Domains.AuthorizationDomain) (outDom V1Domains.AuthorizationDomain, statusCode int, err error) {
	err = authUC.repo.Update(ctx, inDom)
	if err != nil {
		return V1Domains.AuthorizationDomain{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}

func (authUC *authorizationUseCase) Delete(ctx context.Context, id int) (statusCode int, err error) {
	err = authUC.repo.Delete(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (authUC *authorizationUseCase) GetAll(ctx context.Context, params paginate.Params) (outDom paginate.Pagination[V1Domains.AuthorizationDomain], statusCode int, err error) {

	outDom, err = authUC.repo.GetAll(ctx, params)
	if err != nil {
		return paginate.Pagination[V1Domains.AuthorizationDomain]{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}
