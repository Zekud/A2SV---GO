# Library Management System

## Overview
This system is a console-based application implemented in Go to manage a library's books and members. It demonstrates the use of structs, interfaces, methods, slices, and maps.

## Features
- Add a new book to the library.
- Remove an existing book.
- Borrow a book.
- Return a book.
- List all available books.
- List all borrowed books by a member.

## Folder Structure
```
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   └── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod
```

## How to Run
1. Clone the repository.
2. Navigate to the `Task3` folder.
3. Run the application:
   ```bash
   go run main.go
   ```
