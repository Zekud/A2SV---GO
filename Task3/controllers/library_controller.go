package controllers

import (
	"fmt"

	"../models"
	"../services"
)

func StartLibraryConsole() {
	library := services.NewLibrary()
	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var id int
			var title, author string
			fmt.Print("Enter Book ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Book Title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter Book Author: ")
			fmt.Scanln(&author)
			library.AddBook(models.Book{ID: id, Title: title, Author: author, Status: "Available"})
		case 2:
			var id int
			fmt.Print("Enter Book ID to Remove: ")
			fmt.Scanln(&id)
			library.RemoveBook(id)
		case 3:
			var bookID, memberID int
			fmt.Print("Enter Book ID to Borrow: ")
			fmt.Scanln(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scanln(&memberID)
			library.BorrowBook(bookID, memberID)
		case 4:
			var bookID, memberID int
			fmt.Print("Enter Book ID to Return: ")
			fmt.Scanln(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scanln(&memberID)
			library.ReturnBook(bookID, memberID)
		case 5:
			books := library.ListAvailableBooks()
			fmt.Println("Available Books:")
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}
		case 6:
			var memberID int
			fmt.Print("Enter Member ID: ")
			fmt.Scanln(&memberID)
			books := library.ListBorrowedBooks(memberID)
			fmt.Println("Borrowed Books:")
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
