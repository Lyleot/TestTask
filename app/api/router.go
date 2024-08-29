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

// SetRouter инициализирует маршрутизатор Gin с маршрутами и middleware.
func SetRouter(config conf.Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())   // Логирование запросов
	router.Use(gin.Recovery()) // Восстановление от паник

	// Создаем обработчик с доступом к базе данных.
	h := &handler{
		DB: config.DB(),
	}

	// Маршруты для пользователей
	userGroup := router.Group("/user")
	{
		userGroup.POST("/", h.createUser)      // Создание пользователя
		userGroup.GET("/:id", h.getUser)       // Получение пользователя по ID
		userGroup.GET("/all", h.getUsers)      // Получение всех пользователей
		userGroup.PATCH("/:id", h.updateUser)  // Обновление пользователя по ID
		userGroup.DELETE("/:id", h.deleteUser) // Удаление пользователя по ID
	}

	// Маршруты для ролей
	roleGroup := router.Group("/role")
	{
		roleGroup.POST("/", h.createRole)      // Создание роли
		roleGroup.GET("/:id", h.getRole)       // Получение роли по ID
		roleGroup.GET("/all", h.getRoles)      // Получение всех ролей
		roleGroup.PATCH("/:id", h.updateRole)  // Обновление роли по ID
		roleGroup.DELETE("/:id", h.deleteRole) // Удаление роли по ID
	}

	// Маршрут для документации Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
