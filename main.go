package main

import (
	"fmt"
	"log"

	"github.com/aloneen/assign2-ex2/initializers"
	"github.com/aloneen/assign2-ex2/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDatabase()
}

func main() {

	//  Insert users
	initializers.DB.Create(&models.User{Name: "Dias", Age: 22})
	initializers.DB.Create(&models.User{Name: "Aigerim", Age: 25})
	initializers.DB.Create(&models.User{Name: "Rustem", Age: 30})
	fmt.Println("Inserted sample users")

	//  Query users
	var users []models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		log.Fatal("❌ Query failed:", result.Error)
	}

	fmt.Println("Users in DB: ")
	for _, u := range users {
		fmt.Printf("ID: %d | Name: %s | Age: %d\n", u.ID, u.Name, u.Age)
	}

}
