Features
- User Registration and Login with JWT authentication
- Swipe Functionality (like and pass with daily quota)
- Premium Subscription (Unlocks features like unlimited swipes, verified badge)
- Database Migrations with GORM
- Unit Testing
Tech Stack
- Go (Golang) - Backend programming language
- GORM - ORM for Go
- MySQL - Database
- JWT - Authentication (JSON Web Tokens)
Project Structure
The directory structure is designed to be modular and maintainable:
/datingapps
├── /cmd
│   └── main.go
├── /internal
│   ├── /auth
│   │   └── auth.go
│   ├── /swipes
│   │   └── swipes.go
│   ├── /premium
│   │   └── premium.go
├── /pkg
│   ├── /database
│   │   ├── db.go
│   │   └── models.go
│   └── /utils
│       └── utils.go
├── /routes
│   ├── /auth
│   │   └── auth_controller.go
│   ├── /swipes
│   │   └── swipe_controller.go
│   └── /premium
│       └── premium_controller.go
├── /tests
│   ├── auth_test.go
├── /migrations
├── go.mod
├── go.sum
└── README.md

Dependencies
- github.com/golang-jwt/jwt/v4: For JWT token generation and validation.
- gorm.io/gorm: ORM for handling database interactions.
- gorm.io/driver/postgres: PostgreSQL driver for GORM.
- gorm.io/driver/mysql: MySQL driver for GORM.
- github.com/joho/godotenv: For loading environment variables from .env.
- github.com/gin-gonic/gin: For handling HTTP routes.

Setup and Installation
1. Clone the Repository
2. Install Dependencies
3. Configure Environment Variables
Create a .env file in the root of your project with the following variables:
DB_HOST=localhost
DB_PORT=3306
DB_USER=gorm
DB_PASSWORD=gorm
DB_NAME=dating_app
4. Database Setup
MySQL:
If you're using MySQL, run the following commands to create the database:
CREATE DATABASE dating_app;
5. Run Migrations
To create the tables and relationships in your database, run the migration setup.
go run cmd/main.go
This will automatically create the necessary tables based on the AutoMigrate() function in the db.go file.

6. Run the Application
Start the Go application:
go run cmd/main.go
7. API Endpoints
POST /signup: User registration
POST /login: User login (JWT token)
POST /swipe: User swipes left or right
GET /profile: Get user profile
POST /premium: Purchase premium feature
8. Unit Testing
Run the unit tests for the models and services:
go test ./tests/auth_test.go

