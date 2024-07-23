package main

import (
	"github.com/testgo/go-sammy/initializers"
	"github.com/testgo/go-sammy/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
