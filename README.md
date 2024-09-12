# Go User Management Microservice

A simple GoLang microservice that handles user management operations. This service provides a RESTful API for creating, retrieving, updating, and deleting users, with PostgreSQL integration, concurrency using goroutines, and robust error handling.

## Features
- **Create, Read, Update, Delete (CRUD) Users**: Basic user management functionality with fields like name, email, and age.
- **Concurrency**: Logs the time taken by each request using goroutines and channels.
- **Database Integration**: Uses PostgreSQL to persist user data with GORM ORM.
- **Testing**: Includes unit tests for both repository and controller layers.
- **Error Handling**: Robust error handling for database and input validation errors.

---

## Project Structure

```bash
go-user-management/
├── cmd/
│   └── main.go                     # Entry point of the application
├── config/
│   └── config.go                   # Database configuration
├── controllers/
│   └── user_controller.go          # User controller with CRUD operations
├── models/
│   └── user.go                     # User model definition
├── repository/
│   └── user_repository.go          # Repository functions for user management
├── routes/
│   └── user_routes.go              # Routes definition
├── tests/
│   ├── repository_tests/           # Repository unit tests
│   └── controller_tests/           # Controller unit tests
├── utils/
│   └── logger.go                   # Utility functions for logging
├── Dockerfile                      # Dockerfile for building the app
├── docker-compose.yml              # Docker Compose configuration for the app and PostgreSQL
├── go.mod                          # Go module definition
├── go.sum                          # Go dependencies
└── README.md                       # This readme file
```

## Prerequisites
- Go version 1.18+
- Docker compose or PostgreSQL

## Setup Instructions
# 1. Clone the repository
```bash
git clone https://github.com/nubufi/go-user-management.git
cd go-user-management
```
# 2. Environment Variables
You can configure the environment variables in .env file:
```
DB_HOST= localhost
DB_PORT= 5432
DB_USER= user
DB_PASSWORD= password
DB_NAME= usermanagement
HTTP_PORT= 8080
```

# 3. Deploying PostgreSQL(optional)
If you don't have any postgresql server you can deploy it by running the following command
```bash
docker-compose up --build
```

# 4. Running The Application
```bash
go run cmd/main.go
```

## API Endpoints
# 1. Create a User
**POST /users**

**Request Body (JSON):**
```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 30
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 30
}
```

# 2. Get User by ID
**GET /users/:id**

**Response (200 OK):**

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 30
}
```

# 3. Update User by ID
**PUT /users/:id**

**Request Body (JSON):**

```json
{
  "name": "John Updated",
  "email": "john.updated@example.com",
  "age": 31
}
```

**Response (200 OK):**

```json
{
  "id": 1,
  "name": "John Updated",
  "email": "john.updated@example.com",
  "age": 31
}
```

# 4. Delete User by ID
**DELETE /users/:id**

**Response (200 OK):**

```json
{
  "message": "User deleted successfully"
}
```

## Logging
The application logs the time taken for each request using goroutines and channels. You can see logs printed in the terminal where the application is running.

## Testing
This project includes unit tests for both the repository and controller layers. Tests are structured into different directories under the tests/ folder.

# 1. Run All Tests
```bash
go test ./... -v
```

# 2. Running Specific Tests
- **Repository Tests:**
```bash
go test ./tests/repository_tests/ -v
```
- **Controller Tests:**
```bash
go test ./tests/controller_tests/ -v
```
**Test Structure:**
repository_tests/: Contains unit tests for repository logic (CRUD operations on the database).
controller_tests/: Contains tests for the controller layer (API endpoints).

## Database Migrations
GORM handles automatic migrations for this project. When the application starts, it automatically creates the users table based on the models.User struct if it doesn’t already exist.
