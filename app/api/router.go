package api

import (
	"TestTask/app/conf"
	"TestTask/app/store"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type handler struct {
	DB store.Store
}

func SetRouter(config conf.Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	h := &handler{
		DB: config.DB(),
	}

	userGroup := router.Group("/user")
	{
		userGroup.POST("/", h.createUser)
		userGroup.GET("/:id", h.getUser)
		userGroup.GET("/all", h.getUsers)
		userGroup.PATCH("/:id", h.updateUser)
		userGroup.DELETE("/:id", h.deleteUser)
	}

	roleGroup := router.Group("/role")
	{
		roleGroup.POST("/", h.createRole)
		roleGroup.GET("/:id", h.getRole)
		roleGroup.GET("/all", h.getRoles)
		roleGroup.PATCH("/:id", h.updateRole)
		roleGroup.DELETE("/:id", h.deleteRole)
	}

	// Swagger документация
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
