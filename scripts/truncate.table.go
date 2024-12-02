package main

import (
	"basicAuth/initializers"
	"fmt"
	"log"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}
func main() {
	var err error

	// Disable foreign key checks (for PostgreSQL)
	err = initializers.DB.Exec("SET session_replication_role = 'replica'").Error
	if err != nil {
		log.Println("Error disabling foreign key checks:", err)
		return
	}

	// Truncate the pivot table user_permissions first (many-to-many relation)
	err = initializers.DB.Exec("TRUNCATE TABLE user_permissions CASCADE").Error
	if err != nil {
		fmt.Println("Error truncating user_permissions table:", err)
	}

	// Then truncate the 'permissions' table
	err = initializers.DB.Exec("TRUNCATE TABLE permissions RESTART IDENTITY CASCADE").Error
	if err != nil {
		fmt.Println("Error truncating permissions table:", err)
	}

	// Finally, truncate the 'users' table
	err = initializers.DB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error
	if err != nil {
		fmt.Println("Error truncating users table:", err)
	}

	// Re-enable foreign key checks (for PostgreSQL)
	err = initializers.DB.Exec("SET session_replication_role = 'origin'").Error
	if err != nil {
		log.Println("Error re-enabling foreign key checks:", err)
		return
	}

	// Success message
	if err == nil {
		fmt.Println("All tables truncated successfully.")
	}
}
