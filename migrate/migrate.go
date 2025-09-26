package main

import (
	"github.com/aloneen/assign2-ex2/initializers"
	"github.com/aloneen/assign2-ex2/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
