package main

import (
	// "net/http"
	"time"
	// "github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description`
	DueDate     time.Time `json:"due_date"`
	status      string    `json:"status"`
}

var tasks = []Task{
	{
		ID:          "1",
		Title:       "Buy groceries",
		Description: "Milk, Bread, Eggs",
		DueDate:     time.Now().Add(48 * time.Hour),
		status:      "pending",
	},
	{
		ID:          "2",
		Title:       "Read book",
		Description: "Read 'The Go Programming Language'",
		DueDate:     time.Now().Add(72 * time.Hour),
		status:      "in-progress",
	},
}
