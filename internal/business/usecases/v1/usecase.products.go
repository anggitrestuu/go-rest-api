package v1

import (
	"context"
	"net/http"

	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/pkg/paginate"
)

type productUseCase struct {
	repo V1Domains.ProductRepository
}

func NewProductUseCase(repo V1Domains.ProductRepository) V1Domains.ProductUseCase {
	return &productUseCase{
		repo: repo,
	}
}

func (authUC *productUseCase) Store(ctx context.Context, inDom *V1Domains.ProductDomain) (outDom V1Domains.ProductDomain, statusCode int, err error) {
	outDom, err = authUC.repo.Store(ctx, inDom)
	if err != nil {
		return V1Domains.ProductDomain{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusCreated, nil
}

func (authUC *productUseCase) GetByID(ctx context.Context, id int) (outDom V1Domains.ProductDomain, statusCode int, err error) {
	outDom, err = authUC.repo.GetByID(ctx, id)
	if err != nil {
		return V1Domains.ProductDomain{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}

func (authUC *productUseCase) Update(ctx context.Context, inDom *V1Domains.ProductDomain) (outDom V1Domains.ProductDomain, statusCode int, err error) {
	err = authUC.repo.Update(ctx, inDom)
	if err != nil {
		return V1Domains.ProductDomain{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}

func (authUC *productUseCase) Delete(ctx context.Context, id int) (statusCode int, err error) {
	err = authUC.repo.Delete(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (authUC *productUseCase) GetAll(ctx context.Context, params paginate.Params) (outDom paginate.Pagination[V1Domains.ProductDomain], statusCode int, err error) {

	outDom, err = authUC.repo.GetAll(ctx, params)
	if err != nil {
		return paginate.Pagination[V1Domains.ProductDomain]{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}
