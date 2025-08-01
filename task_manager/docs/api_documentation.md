# Task Management API Documentation

## Overview
This API allows you to manage tasks using a RESTful interface. It now integrates with MongoDB for persistent data storage.

## MongoDB Integration
- The API uses MongoDB as the database to store tasks.
- The MongoDB URI and database name are configured using environment variables:
  - `MONGO_URI`: The connection string for MongoDB (e.g., `mongodb://localhost:27017`).
  - `MONGO_DB`: The name of the database (e.g., `task_manager`).
- Ensure you have a `.env` file in the root of your project with the following content:
  ```env
  MONGO_URI=mongodb://localhost:27017
  MONGO_DB=task_manager
  ```

## JWT Integration
- The API uses JWT for authentication and authorization.
- The secret key for signing and verifying JWTs is configured using an environment variable:
  - `JWT_SECRET_KEY`: The secret key for JWT (e.g., `your_secret_key`).
- Ensure you have a `.env` file in the root of your project with the following content:
  ```env
  JWT_SECRET_KEY=your_secret_key
  ```

## Endpoints

### 1. Get All Tasks
- **URL**: `/tasks`
- **Method**: `GET`
- **Response**:
  ```json
  [
    {
      "id": "1",
      "title": "Buy groceries",
      "description": "Milk, Bread, Eggs",
      "due_date": "2025-07-19T12:00:00Z",
      "status": "pending"
    }
  ]
  ```

### 2. Get Task by ID
- **URL**: `/tasks/:id`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "id": "1",
    "title": "Buy groceries",
    "description": "Milk, Bread, Eggs",
    "due_date": "2025-07-19T12:00:00Z",
    "status": "pending"
  }
  ```

### 3. Create Task
- **URL**: `/tasks`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "title": "New Task",
    "description": "Task description",
    "due_date": "2025-07-20T12:00:00Z",
    "status": "pending"
  }
  ```
- **Response**:
  ```json
  {
    "id": "2",
    "title": "New Task",
    "description": "Task description",
    "due_date": "2025-07-20T12:00:00Z",
    "status": "pending"
  }
  ```

### 4. Update Task
- **URL**: `/tasks/:id`
- **Method**: `PUT`
- **Request Body**:
  ```json
  {
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2025-07-21T12:00:00Z",
    "status": "completed"
  }
  ```
- **Response**:
  ```json
  {
    "id": "1",
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2025-07-21T12:00:00Z",
    "status": "completed"
  }
  ```

### 5. Delete Task
- **URL**: `/tasks/:id`
- **Method**: `DELETE`
- **Response**:
  ```json
  {}
  ```

## Setup Instructions
1. Ensure MongoDB is running locally or use a cloud-hosted MongoDB instance (e.g., MongoDB Atlas).
2. Create a `.env` file in the root of your project with the MongoDB URI and database name.
3. Run the application:
   ```bash
   go run main.go
   ```
4. Use tools like Postman or curl to test the API endpoints.

## Clean Architecture Structure

The codebase is organized using Clean Architecture principles for better maintainability and testability. The main layers are:

- **Domain/**: Core business entities and logic (e.g., Task, User).
- **Usecases/**: Application-specific business rules (e.g., task and user use cases).
- **Repositories/**: Interfaces and implementations for data access (e.g., MongoDB repositories).
- **Infrastructure/**: External services and middleware (e.g., JWT, password hashing, authentication middleware).
- **Delivery/**: HTTP server, controllers, and routers (handles incoming requests and responses).

### Entry Point
- The application entry point is now `Delivery/main.go`.

### Contribution Guidelines
- Add new business logic to the appropriate use case in `Usecases/`.
- Add new data access logic by extending repository interfaces and implementations in `Repositories/`.
- Add new endpoints by updating controllers and routers in `Delivery/`.
- Keep domain models in `Domain/` free from external dependencies.
