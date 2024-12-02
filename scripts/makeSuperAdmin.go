package main

import (
	"basicAuth/initializers"
	"basicAuth/models"
	"fmt"
	"os"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	var userName string
	var password string
	var isAdmin bool

	var user models.User
	initializers.DB.Where("is_admin=?", true).First(&user)
	if user.ID != 0 {
		fmt.Print("Already have A Admin")
		return
	}

	// Prompt for username
	fmt.Print("Enter your username: ")
	_, err := fmt.Scanln(&userName)
	if err != nil {
		fmt.Println("Error reading username:", err)
		os.Exit(1)
	}

	// Prompt for password
	fmt.Print("Enter your password: ")
	_, err = fmt.Scanln(&password)
	if err != nil {
		fmt.Println("Error reading password:", err)
		os.Exit(1)
	}

	var admin models.User
	admin.UserName = userName
	admin.Password = password
	admin.IsAdmin = true
	initializers.DB.Create(&admin)
	fmt.Printf("\nUser Details:\nUsername: %s\nPassword: %s\nAdmin: %t\n", userName, password, isAdmin)
}
