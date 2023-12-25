package v1

import (
	"net/http"

	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/internal/http/datatransfers/requests"
	"github.com/anggitrestuu/go-rest-api/internal/http/datatransfers/responses"
	"github.com/anggitrestuu/go-rest-api/internal/utils"
	"github.com/anggitrestuu/go-rest-api/pkg/validators"
	"github.com/gin-gonic/gin"
)

type AuthorizationHandler struct {
	useCase V1Domains.AuthorizationUseCase
}

func NewAuthorizationHandler(useCase V1Domains.AuthorizationUseCase) AuthorizationHandler {
	return AuthorizationHandler{
		useCase: useCase,
	}
}

// @Summary Create new authorization
// @Description Create new authorization
// @Tags authorization
// @Accept json
// @Produce json
// @Param authorization body requests.AuthorizationRequest true "Create new authorization"
// @Success 201 {object} map[string]interface{} "create new authorization success"
// @Router /api/v1/authorizations [post]
func (authH AuthorizationHandler) Store(ctx *gin.Context) {
	var AuthorizationRequest requests.AuthorizationRequest
	if err := ctx.ShouldBindJSON(&AuthorizationRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(AuthorizationRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	inDomain := AuthorizationRequest.ToV1Domain()
	outDomain, statusCode, err := authH.useCase.Store(ctx.Request.Context(), inDomain)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "create new authorization success", responses.FromAuthorizationV1Domain(outDomain))

}

// @Summary Get authorization by id
// @Description Get authorization by id
// @Tags authorization
// @Accept json
// @Produce json
// @Param id path int true "authorization id"
// @Success 200 {object} map[string]interface{} "get authorization by id success"
// @Router /api/v1/authorizations/{id} [get]
func (authH AuthorizationHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	outDomain, statusCode, err := authH.useCase.GetByID(ctx.Request.Context(), utils.StringToInt(id))
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "get authorization by id success", responses.FromAuthorizationV1Domain(outDomain))
}
