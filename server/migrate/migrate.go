package main

import (
	"github.com/trishan9/Gin-CRUD/initializers"
	"github.com/trishan9/Gin-CRUD/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// initializers.DB.Migrator().DropTable(&models.User{})
	// initializers.DB.Migrator().DropTable(&models.Post{})
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
}
