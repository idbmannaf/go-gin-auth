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

	var err error
	// Drop 'permissions' table
	err = initializers.DB.Migrator().DropTable(&models.Permission{})
	if err != nil {
		fmt.Println("Error Deleting Permission Table:", err)
		return
	}
	fmt.Println("Permission table dropped successfully.")

	// Drop 'user_permissions' table
	err = initializers.DB.Migrator().DropTable("user_permissions")
	if err != nil {
		fmt.Println("Error Deleting user_permissions Table:", err)
		return
	}
	fmt.Println("user_permissions table dropped successfully.")

	// Drop 'users' table
	err = initializers.DB.Migrator().DropTable(&models.User{})
	if err != nil {
		fmt.Println("Error Deleting Users Table:", err)
		return
	}
	fmt.Println("Users table dropped successfully.")
}
