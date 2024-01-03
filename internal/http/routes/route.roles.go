package routes

import (
	V1Usecase "github.com/anggitrestuu/go-rest-api/internal/business/usecases/v1"
	V1PostgresRepository "github.com/anggitrestuu/go-rest-api/internal/datasources/repositories/postgres/v1"
	V1Handler "github.com/anggitrestuu/go-rest-api/internal/http/handlers/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type rolesRoutes struct {
	V1Handler      V1Handler.RoleHandler
	router         *gin.RouterGroup
	db             *gorm.DB
	authMiddleware gin.HandlerFunc
}

func NewRolesRoute(router *gin.RouterGroup, db *gorm.DB, authMiddleware gin.HandlerFunc) *rolesRoutes {
	V1RoleRepository := V1PostgresRepository.NewRoleRepository(db)
	V1RoleUsecase := V1Usecase.NewRoleUseCase(V1RoleRepository)
	V1RoleHandler := V1Handler.NewRoleHandler(V1RoleUsecase)

	return &rolesRoutes{V1Handler: V1RoleHandler, router: router, db: db,
		authMiddleware: authMiddleware,
	}

}

func (r *rolesRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		// roles
		roleRoute := V1Route.Group("/roles")
		// roleRoute.Use(r.authMiddleware)
		{
			roleRoute.POST("", r.V1Handler.Store)
			roleRoute.PUT("/:id", r.V1Handler.Update)
			roleRoute.GET("/:id", r.V1Handler.GetByID)
			roleRoute.GET("/roles", r.V1Handler.GetAll)
			roleRoute.DELETE("/:id", r.V1Handler.Delete)

		}
	}

}
