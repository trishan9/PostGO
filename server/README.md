# Gin CRUD API

This is a CRUD API built using the Gin framework for creating, reading, updating, and deleting posts. The API also includes user authentication and authorization features.

## Features

- User authentication (sign up, login, email verification)
- CRUD operations for posts
- Swagger documentation for API endpoints

## Technologies Used

- Go programming language
- Gin framework
- GORM (Go Object Relational Mapper) for database interactions
- Swagger for API documentation
- JWT for authorization
- Cloudinary for file uploads
- Bcrypt for hashing
- Sendgrid for emails

## Installation

1. Clone the repository:

```bash
git clone git@github.com:trishan9/Gin-CRUD.git
cd Gin-CRUD
```

2. Install dependencies

```go
go mod tidy
```

3. Setup environment variables

```bash
cp .env.sample .env
```

4. Run the application:

```bash
go run main.go
```

## API Documentation

The API documentation is generated using Swagger. You can access it at `/api/docs/index.html` route after starting the application.

## Endpoints

- POST `/api/auth/signup`: Sign up for a new account.
- POST `/api/auth/login`: Log in to an existing account.
- POST `/api/auth/signup/validate`: Validate email address using OTP.
- POST `/api/auth/signup/regenerate`: Regenerate OTP for email verification (if needed).
- POST `/api/posts`: Create a new post.
- GET `/api/posts`: Get all posts.
- GET `/api/posts/:id`: Get a post by ID.
- PATCH `/api/posts/:id`: Update a post.
- DELETE `/api/posts/:id`: Delete a post.
