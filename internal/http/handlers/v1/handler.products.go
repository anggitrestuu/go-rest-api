package v1

import (
	"net/http"

	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/internal/http/datatransfers/requests"
	"github.com/anggitrestuu/go-rest-api/internal/http/datatransfers/responses"
	"github.com/anggitrestuu/go-rest-api/internal/utils"
	"github.com/anggitrestuu/go-rest-api/pkg/paginate"
	"github.com/anggitrestuu/go-rest-api/pkg/validators"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	useCase V1Domains.ProductUseCase
}

func NewProductHandler(useCase V1Domains.ProductUseCase) ProductHandler {
	return ProductHandler{
		useCase: useCase,
	}
}

// @Summary Create new product
// @Description Create new product
// @Tags product
// @Accept json
// @Produce json
// @Param product body requests.ProductRequest true "Create new product"
// @Success 201 {object} map[string]interface{} "create new product success"
// @Router /api/v1/products [post]
func (h ProductHandler) Store(ctx *gin.Context) {
	var ProductRequest requests.ProductRequest
	if err := ctx.ShouldBindJSON(&ProductRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(ProductRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	inDomain := ProductRequest.ToV1Domain()
	outDomain, statusCode, err := h.useCase.Store(ctx.Request.Context(), inDomain)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "create new product success", responses.FromProductV1Domain(outDomain))

}

// @Summary Get product by id
// @Description Get product by id
// @Tags product
// @Accept json
// @Produce json
// @Param id path int true "product id"
// @Success 200 {object} map[string]interface{} "get product by id success"
// @Router /api/v1/products/{id} [get]
func (h ProductHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	outDomain, statusCode, err := h.useCase.GetByID(ctx.Request.Context(), utils.StringToInt(id))
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "get product by id success", responses.FromProductV1Domain(outDomain))
}

// @Summary delete product by id
// @Description delete product by id
// @Tags product
// @Accept json
// @Produce json
// @Param id path int true "product id"
// @Success 200 {object} map[string]interface{} "Delete product by id success"
// @Router /api/v1/products/{id} [delete]
func (h ProductHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	statusCode, err := h.useCase.Delete(ctx.Request.Context(), utils.StringToInt(id))
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "delete product by id success", nil)
}

// @Summary Get all product
// @Description Get all product
// @Tags product
// @Accept json
// @Produce json
// @Param limit query string false "Limit" default(10)
// @Param page query string false "Page" default(1)
// @Param sort_by query string false "Sort By"
// @Param filters query string false "Filters"
// @Success 200 {object} map[string]interface{} "get all product success"
// @Router /api/v1/products [get]
func (h ProductHandler) GetAll(ctx *gin.Context) {
	var queryParams paginate.Params
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Invalid query parameters")
		return
	}

	outDomain, statusCode, err := h.useCase.GetAll(ctx.Request.Context(), queryParams)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	roleResponsePagination := responses.TransformPagination(outDomain, responses.FromProductV1Domain)

	NewSuccessResponse(ctx, statusCode, "get all product success", roleResponsePagination)
}

// @Summary Update product by id
// @Description Update product by id
// @Tags product
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Param product body requests.ProductRequest true "Update new product"
// @Success 201 {object} map[string]interface{} "update product success"
// @Router /api/v1/products/{id} [put]
func (h ProductHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var ProductRequest requests.ProductRequest
	if err := ctx.ShouldBindJSON(&ProductRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(ProductRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	inDomain := ProductRequest.ToV1Domain()
	inDomain.ID = utils.StringToInt(id)

	outDomain, statusCode, err := h.useCase.Update(ctx.Request.Context(), inDomain)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "update product by id success", responses.FromProductV1Domain(outDomain))
}
