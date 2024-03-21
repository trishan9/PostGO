package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/trishan9/Gin-CRUD/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	port := os.Getenv("PORT")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello GO CRUD",
		})
	})

	r.Run(port) // 8080 port by default
}
