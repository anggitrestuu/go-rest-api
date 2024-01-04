package routes

import (
	V1Usecase "github.com/anggitrestuu/go-rest-api/internal/business/usecases/v1"
	V1PostgresRepository "github.com/anggitrestuu/go-rest-api/internal/datasources/repositories/postgres/v1"
	V1Handlers "github.com/anggitrestuu/go-rest-api/internal/http/handlers/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authorizationsRoutes struct {
	V1Handler      V1Handlers.AuthorizationHandler
	router         *gin.RouterGroup
	db             *gorm.DB
	authMiddleware gin.HandlerFunc
}

func NewAuthorizationsRoute(router *gin.RouterGroup, db *gorm.DB, authMiddleware gin.HandlerFunc) *authorizationsRoutes {
	V1AuthRepository := V1PostgresRepository.NewAuthorizationRepository(db)
	V1AuthUseCase := V1Usecase.NewAuthorizationUseCase(V1AuthRepository)
	V1AuthHandler := V1Handlers.NewAuthorizationHandler(V1AuthUseCase)

	return &authorizationsRoutes{V1Handler: V1AuthHandler, router: router, db: db,
		authMiddleware: authMiddleware,
	}
}

func (r *authorizationsRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		// authorizations
		authorizationRoute := V1Route.Group("/authorizations")
		{
			authorizationRoute.GET("/", r.V1Handler.GetAll)
			authorizationRoute.POST("/", r.V1Handler.Store)
			authorizationRoute.GET("/:id", r.V1Handler.GetByID)
			authorizationRoute.PUT("/:id", r.V1Handler.Update)
			authorizationRoute.DELETE("/:id", r.V1Handler.Delete)
		}
	}
}
