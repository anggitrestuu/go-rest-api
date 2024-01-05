package v1

import (
	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/internal/http/datatransfers/responses"
	"github.com/anggitrestuu/go-rest-api/internal/utils"
	"github.com/gin-gonic/gin"
)

type RoleAuthorizationHandler struct {
	useCase V1Domains.RoleAuthorizationUseCase
}

func NewRoleAuthorizationHandler(useCase V1Domains.RoleAuthorizationUseCase) RoleAuthorizationHandler {
	return RoleAuthorizationHandler{
		useCase: useCase,
	}
}

// @Summary Assign authorization to role
// @Description Assign authorization to role
// @Tags role-authorization
// @Accept json
// @Produce json
// @Param roles_id path int true "Role ID"
// @Param authorizations_id path int true "Authorization ID"
// @Success 200 {object} map[string]interface{} "assign authorization to role success"
// @Router /api/v1/role-authorizations/{roles_id}/{authorizations_id} [post]
func (roleAuthH RoleAuthorizationHandler) AssignAuthorizationToRole(ctx *gin.Context) {
	roleID := utils.StringToInt(ctx.Param("roles_id"))
	authorizationID := utils.StringToInt(ctx.Param("authorizations_id"))

	outDomain, statusCode, err := roleAuthH.useCase.AssignAuthorizationToRole(ctx.Request.Context(), roleID, authorizationID)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "assign authorization to role success", responses.FromRoleAuthorizationV1Domain(outDomain))
}

// @Summary Remove authorization from role
// @Description Remove authorization from role
// @Tags role-authorization
// @Accept json
// @Produce json
// @Param roles_id path int true "Role ID"
// @Param authorizations_id path int true "Authorization ID"
// @Success 200 {object} map[string]interface{} "remove authorization from role success"
// @Router /api/v1/role-authorizations/{roles_id}/{authorizations_id} [delete]
func (roleAuthH RoleAuthorizationHandler) RemoveAuthorizationFromRole(ctx *gin.Context) {
	roleID := utils.StringToInt(ctx.Param("roles_id"))
	authorizationID := utils.StringToInt(ctx.Param("authorizations_id"))

	outDomain, statusCode, err := roleAuthH.useCase.RemoveAuthorizationFromRole(ctx.Request.Context(), roleID, authorizationID)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "remove authorization from role success", responses.FromRoleAuthorizationV1Domain(outDomain))
}

// @Summary Get authorizations by role id
// @Description Get authorizations by role id
// @Tags role-authorization
// @Accept json
// @Produce json
// @Param roles_id path int true "Role ID"
// @Success 200 {object} map[string]interface{} "get authorizations by role id success"
// @Router /api/v1/role-authorizations/{roles_id} [get]
func (roleAuthH RoleAuthorizationHandler) GetAuthorizationsByRoleID(ctx *gin.Context) {
	roleID := utils.StringToInt(ctx.Param("roles_id"))

	outDomains, statusCode, err := roleAuthH.useCase.GetAuthorizationsByRoleID(ctx.Request.Context(), roleID)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "get authorizations by role id success", responses.FromAuthorizationV1Domains(outDomains))
}
