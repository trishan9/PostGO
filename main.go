package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/trishan9/Gin-CRUD/controllers"
	_ "github.com/trishan9/Gin-CRUD/docs"
	"github.com/trishan9/Gin-CRUD/initializers"
	"github.com/trishan9/Gin-CRUD/middlewares"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// @title Gin PostsApp API
// @version 1.0
// @description This is API for PostsApp using Gin

// @host localhost:3000
// @BasePath /api
func main() {
	port := os.Getenv("PORT")

	r := gin.Default()

	r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Gin CRUD API",
		})
	})

	baseRouter := r.Group("/api")
	// Auth Routes
	authRouter := baseRouter.Group("/auth")
	authRouter.POST("/signup", controllers.SignUp)
	authRouter.POST("/signup/validate", controllers.ValidateOtp)
	authRouter.POST("/signup/regenerate", controllers.RegenerateOtp)
	authRouter.POST("/login", controllers.Login)

	// Posts Routes
	postsRouter := baseRouter.Group("/posts")
	postsRouter.POST("/", middlewares.IsAuthenticated, controllers.CreatePost)
	postsRouter.GET("/", middlewares.IsAuthenticated, controllers.GetPosts)
	postsRouter.GET("/:id", middlewares.IsAuthenticated, controllers.GetPostById)
	postsRouter.PATCH("/:id", middlewares.IsAuthenticated, controllers.UpdatePost)
	postsRouter.DELETE("/:id", middlewares.IsAuthenticated, controllers.DeletePost)

	r.Run(":" + port) // 8080 port by default
}
