package routes

import (
	V1Usecase "github.com/anggitrestuu/go-rest-api/internal/business/usecases/v1"
	V1PostgresRepository "github.com/anggitrestuu/go-rest-api/internal/datasources/repositories/postgres/v1"
	V1Handlers "github.com/anggitrestuu/go-rest-api/internal/http/handlers/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productsRoutes struct {
	V1Handler      V1Handlers.ProductHandler
	router         *gin.RouterGroup
	db             *gorm.DB
	authMiddleware gin.HandlerFunc
}

func NewProductsRoute(router *gin.RouterGroup, db *gorm.DB, authMiddleware gin.HandlerFunc) *productsRoutes {
	V1AuthRepository := V1PostgresRepository.NewProductRepository(db)
	V1AuthUseCase := V1Usecase.NewProductUseCase(V1AuthRepository)
	V1AuthHandler := V1Handlers.NewProductHandler(V1AuthUseCase)

	return &productsRoutes{V1Handler: V1AuthHandler, router: router, db: db,
		authMiddleware: authMiddleware,
	}
}

func (r *productsRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		// products
		productRoute := V1Route.Group("/products")
		{
			productRoute.GET("/", r.V1Handler.GetAll)
			productRoute.POST("/", r.V1Handler.Store)
			productRoute.GET("/:id", r.V1Handler.GetByID)
			productRoute.PUT("/:id", r.V1Handler.Update)
			productRoute.DELETE("/:id", r.V1Handler.Delete)
		}
	}
}
