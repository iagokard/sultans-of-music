package routes

import (
	"net/http"

	_ "som/docs"
	"som/internal/config"
	"som/internal/middleware"

	artistHandler "som/internal/handlers/artist"
	productHandler "som/internal/handlers/product"
	saleHandler "som/internal/handlers/sale"
	userHandler "som/internal/handlers/user"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {
	cfg := config.Load()

	router.Use(middleware.TimeoutMiddleware())
	router.Use(middleware.CORS())

	router.GET("/apidoc", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/apidoc/index.html")
	})
	router.GET("/apidoc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")

	{
		artist := api.Group("/artist")

		artist.POST("/create", artistHandler.CreateArtist)

		artist.Use(middleware.JWT(cfg.JWTSecret))
	}

	{
		product := api.Group("/product")

		product.Use(middleware.JWT(cfg.JWTSecret))

		product.POST("/create", productHandler.CreateProduct)
		product.POST("/list", productHandler.GetProductListPage)
		product.POST("/sell", productHandler.SellProduct)
		product.PUT("/update", productHandler.UpdateProduct)
		product.GET("/all", productHandler.GetAllProducts)
		product.GET("/:id", productHandler.GetProductByID)
	}

	{
		inventory := api.Group("/inventory")

		inventory.Use(middleware.JWT(cfg.JWTSecret))

		inventory.PUT("/update", productHandler.UpdateInventory)
	}

	{
		user := api.Group("/user")

		user.POST("/login", userHandler.LoginUser)
		user.POST("/create", userHandler.CreateUser)

		user.Use(middleware.JWT(cfg.JWTSecret))

		user.GET("/:id", userHandler.GetUserByID)
	}

	{
		sale := api.Group("/sale")

		sale.Use(middleware.JWT(cfg.JWTSecret))

		sale.GET("/:id", saleHandler.GetSaleByID)
		sale.GET("/all", saleHandler.GetAllSales)
	}
}
