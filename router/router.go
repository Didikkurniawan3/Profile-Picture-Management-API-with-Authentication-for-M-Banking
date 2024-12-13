package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/controllers"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/database"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/middlewares"
	"net/http"
)

func RouteInit() *gin.Engine {
	// Inisialisasi router Gin
	r := gin.Default()

	// Menyediakan file statis untuk gambar
	r.Static("/images", "./static/images")

	// Menyediakan favicon.ico
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	// Mengambil instance database
	db := database.GetDB()

	// Membuat controller
	userController := controllers.NewUserController(db)
	photoController := controllers.NewPhotoController(db)

	// Grup API v1
	api := r.Group("/api/v1")

	// Rute untuk users
	userRoutes := api.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.PUT("/:userId", userController.Update)
		userRoutes.DELETE("/:userId", userController.Delete)
	}

	// Rute untuk photo dengan middleware autentikasi
	photoRoutes := api.Group("/photo")
	photoRoutes.Use(middlewares.AuthMiddleware(db))
	{
		photoRoutes.GET("/", photoController.Get)
		photoRoutes.POST("/", photoController.Create)
		photoRoutes.PUT("/", photoController.Update)
		photoRoutes.DELETE("/", photoController.Delete)
	}

	// Menambahkan rute untuk root ("/") yang mengembalikan respons sederhana
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the API!",
		})
	})

	// Kembalikan router
	return r
}
