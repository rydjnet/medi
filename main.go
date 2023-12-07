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
	router.GET("/medi/:name", controllers.FindRow)
	router.PATCH("/medi/:name", controllers.UpdateRow)
	router.DELETE("/medi/:name", controllers.DelRow)
	router.Run("localhost:8080")
}
