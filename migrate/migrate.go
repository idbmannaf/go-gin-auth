package main

import (
	"basicAuth/initializers"
	"basicAuth/models"
	"fmt"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Permission{})
	fmt.Println("Migration Successfully Done:")
}
