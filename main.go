package main

import (
	"medi/controllers"
	"medi/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	models.ConnectDB()
	defer models.DB.Close()
	router.POST("/medi", controllers.CreatePost)
	router.GET("/medi", controllers.GetList)
	router.Run("localhost:8080")
}
