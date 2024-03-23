package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/trishan9/Gin-CRUD/controllers"
	"github.com/trishan9/Gin-CRUD/initializers"
	"github.com/trishan9/Gin-CRUD/middlewares"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	port := os.Getenv("PORT")

	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Auth Routes
	r.POST("/api/auth/signup", controllers.SignUp)
	r.POST("/api/auth/signup/validate", controllers.ValidateOtp)
	r.POST("/api/auth/signup/regenerate", controllers.RegenerateOtp)
	r.POST("/api/auth/login", controllers.Login)

	// Posts Routes
	r.POST("/api/posts", middlewares.IsAuthenticated, controllers.CreatePosts)
	r.GET("/api/posts", middlewares.IsAuthenticated, controllers.GetPosts)
	r.GET("/api/posts/:id", middlewares.IsAuthenticated, controllers.GetPostById)
	r.PATCH("/api/posts/:id", middlewares.IsAuthenticated, controllers.UpdatePost)
	r.DELETE("/api/posts/:id", middlewares.IsAuthenticated, controllers.DeletePost)

	r.Run(port) // 8080 port by default
}
