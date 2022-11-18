package main

import (
	"api_gin_gorm/DB"
	"api_gin_gorm/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	DB.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello SalikhovR!"})
	})
	r.GET("/books", server.FindBooks)
	r.POST("/books", server.CreateBook)
	r.GET("/books/:id", server.FindBook)
	r.PATCH("/books/:id", server.UpdateBook)
	r.DELETE("/books/:id", server.DeleteBook)

	r.Run("localhost:8080")
}
