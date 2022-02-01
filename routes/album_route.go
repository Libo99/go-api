package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(router *gin.Engine) {
	router.GET("/albums", controllers.GetAlbums)
	router.GET("albums/:id", controllers.GetById)
	router.POST("/albums", controllers.PostAlbums)
}
