package main

import (
	"task_manager/router"
)

func main() {
	r := router.SetupRouter()
	r.Run() // Start the server on the default port 8080
}
