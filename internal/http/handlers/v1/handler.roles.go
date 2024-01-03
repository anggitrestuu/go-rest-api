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

type RoleHandler struct {
	useCase V1Domains.RoleUseCase
}

func NewRoleHandler(useCase V1Domains.RoleUseCase) RoleHandler {
	return RoleHandler{
		useCase: useCase,
	}
}

// @Summary Create new role
// @Description Create new role
// @Tags role
// @Accept json
// @Produce json
// @Param role body requests.RoleRequest true "Create new role"
// @Success 201 {object} responses.RoleResponse{}
// @Router /api/v1/roles [post]
func (h RoleHandler) Store(ctx *gin.Context) {
	var RoleRequest requests.RoleRequest
	if err := ctx.ShouldBindJSON(&RoleRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(RoleRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	inDomain := RoleRequest.ToV1Domain()
	outDomain, statusCode, err := h.useCase.Store(ctx.Request.Context(), inDomain)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "create new role success", responses.FromRoleV1Domain(outDomain))
}

// @Summary Get role by id
// @Description Get role by id
// @Tags role
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} responses.RoleResponse{}
// @Router /api/v1/roles/{id} [get]
func (h RoleHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	outDomain, statusCode, err := h.useCase.GetByID(ctx.Request.Context(), utils.StringToInt(id))
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "get role by id success", responses.FromRoleV1Domain(outDomain))
}

// @Summary Update role by id
// @Description Update role by id
// @Tags role
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Param role body requests.RoleRequest true "Update role"
// @Success 200 {object} responses.RoleResponse{}
// @Router /api/v1/roles/{id} [put]
func (h RoleHandler) Update(ctx *gin.Context) {
	var RoleRequest requests.RoleRequest
	if err := ctx.ShouldBindJSON(&RoleRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(RoleRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	inDomain := RoleRequest.ToV1Domain()
	outDomain, statusCode, err := h.useCase.Update(ctx.Request.Context(), inDomain)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "update role success", responses.FromRoleV1Domain(outDomain))
}

// @Summary Delete role by id
// @Description Delete role by id
// @Tags role
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {string} string "delete role success"
// @Router /api/v1/roles/{id} [delete]
func (h RoleHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	statusCode, err := h.useCase.Delete(ctx.Request.Context(), utils.StringToInt(id))
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "delete role success", nil)
}

// @Summary Get all role
// @Description Get all role
// @Tags role
// @Accept json
// @Produce json
// @Success 200 {object} []responses.RoleResponse{}
// @Router /api/v1/roles [get]
func (h RoleHandler) GetAll(ctx *gin.Context) {
	outDomain, statusCode, err := h.useCase.GetAll(ctx.Request.Context())
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "get all role success", outDomain)
}
