package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/controllers"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/database"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/middlewares"
)

func RouteInit() *gin.Engine {
	r := gin.Default()
	r.Static("/images", "./static/images")

	db := database.GetDB()

	userController := controllers.NewUserController(db)
	photoController := controllers.NewPhotoController(db)

	api := r.Group("/api/v1")

	userRoutes := api.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.PUT("/:userId", userController.Update)
		userRoutes.DELETE("/:userId", userController.Delete)
	}

	photoRoutes := api.Group("/photo")
	photoRoutes.Use(middlewares.AuthMiddleware(db))
	{
		photoRoutes.GET("/", photoController.Get)
		photoRoutes.POST("/", photoController.Create)
		photoRoutes.PUT("/", photoController.Update)
		photoRoutes.DELETE("/", photoController.Delete)
	}

	return r
}
