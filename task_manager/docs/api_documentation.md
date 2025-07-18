# Task Management API Documentation

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
