package routes

import (
	V1Usecase "github.com/anggitrestuu/go-rest-api/internal/business/usecases/v1"
	V1PostgresRepository "github.com/anggitrestuu/go-rest-api/internal/datasources/repositories/postgres/v1"
	V1Handlers "github.com/anggitrestuu/go-rest-api/internal/http/handlers/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// file : internal/http/routes/router.role-authorizations.go

type roleAuthorizationsRoutes struct {
	V1Handler      V1Handlers.RoleAuthorizationHandler
	router         *gin.RouterGroup
	db             *gorm.DB
	authMiddleware gin.HandlerFunc
}

func NewRoleAuthorizationsRoute(router *gin.RouterGroup, db *gorm.DB, authMiddleware gin.HandlerFunc) *roleAuthorizationsRoutes {
	V1RoleAuthRepository := V1PostgresRepository.NewRoleAuthorizationRepository(db)
	V1RoleAuthUseCase := V1Usecase.NewRoleAuthorizationUseCase(V1RoleAuthRepository)
	V1RoleAuthHandler := V1Handlers.NewRoleAuthorizationHandler(V1RoleAuthUseCase)

	return &roleAuthorizationsRoutes{V1Handler: V1RoleAuthHandler, router: router, db: db,
		authMiddleware: authMiddleware,
	}

}

func (r *roleAuthorizationsRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		// roles
		roleAuthorizationRoute := V1Route.Group("/role-authorizations")
		{
			roleAuthorizationRoute.POST("/:roles_id/:authorizations_id", r.V1Handler.AssignAuthorizationToRole)
			roleAuthorizationRoute.DELETE("/:roles_id/:authorizations_id", r.V1Handler.RemoveAuthorizationFromRole)
			roleAuthorizationRoute.GET("/:roles_id", r.V1Handler.GetAuthorizationsByRoleID)
		}
	}
}
