package main

import (
	"task_manager/config"
	"task_manager/router"
)

func main() {
	// Connect to MongoDB
	config.ConnectMongoDB()

	// Set up routes
	r := router.SetupRouter()

	// Start the server
	r.Run(":8080")
}
