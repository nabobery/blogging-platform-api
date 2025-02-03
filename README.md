# Blogging Platform API

https://roadmap.sh/projects/blogging-platform-api

This is a simple RESTful API for a personal blogging platform. The API allows users to perform basic CRUD operations—Create, Read, Update, Delete—on blog posts. It is designed to help you understand RESTful conventions, HTTP methods, error handling, and working with databases such as PostgreSQL.

## Goals

The purpose of this project is to:

- Understand what RESTful APIs are and learn best practices.
- Learn how to create, read, update, and delete resources using HTTP methods (GET, POST, PUT, DELETE).
- Gain experience with status codes and proper error handling.
- Learn to work with a relational database (PostgreSQL) using an ORM (GORM).
- Understand dependency injection and environment-based configurations.
- Compare different Go web frameworks (why we choose Gorilla Mux over Gin for this project).

## Features

- **Create a Blog Post:**
  `POST /posts`
  Creates a new blog post with the following JSON body:

  ```json
  {
    "title": "My First Blog Post",
    "content": "This is the content of my first blog post.",
    "category": "Technology",
    "tags": ["Tech", "Programming"]
  }
  ```

  Returns a `201 Created` response with the created post or a `400 Bad Request` for invalid input.

- **Get a Single Blog Post:**
  `GET /posts/{id}`
  Returns a specific blog post with status `200 OK` or `404 Not Found` if the post does not exist.

- **Get All Blog Posts (with optional filtering):**
  `GET /posts`
  Optional query parameter `term` allows for filtering blog posts (wildcard search on title, content, or category).

- **Update a Blog Post:**
  `PUT /posts/{id}`
  Updates an existing post. Expects a JSON payload similar to the create endpoint.
  Returns a `200 OK` response with the updated post, or `400 Bad Request`/`404 Not Found` if validation fails or the post doesn’t exist.

- **Delete a Blog Post:**
  `DELETE /posts/{id}`
  Deletes an existing blog post. Returns `204 No Content` if the deletion is successful or `404 Not Found` if the post cannot be found.

## Prerequisites

- [Go 1.21+](https://golang.org/dl/)
- [PostgreSQL](https://www.postgresql.org/) (or any PostgreSQL cloud provider like Neon)
- (Optional) Docker for running PostgreSQL locally

## Environment Variables

Create a `.env` file in the project root to manage configuration:

```env
# PostgreSQL connection string (update with your credentials)
DATABASE_URL="postgres://username:password@localhost:5432/blogdb?sslmode=disable"

# The port the API will listen on
PORT="8080"
```

> **Note:** Update the `DATABASE_URL` string as per your database settings.

## Installation & Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/nabobery/blogging-platform-api.git
   cd blogging-platform-api
   ```

2. **Install Dependencies**

   Use Go modules to install the required dependencies:

   ```bash
   go mod tidy
   ```

3. **Database Setup**

   Create a PostgreSQL database (or use an existing one) and update the `DATABASE_URL` in the `.env` file.

4. **Run the Application**

   Start the server by executing:

   ```bash
   go run main.go
   ```

   You should see logs indicating a successful connection to the database and that the server is running (e.g., "Server is running on port 8080").

## API Usage

### Example cURL Commands

- **Create a Blog Post:**

  ```bash
  curl --location --request POST 'http://localhost:8080/posts' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "title": "My First Blog Post",
      "content": "This is the content of my first blog post.",
      "category": "Technology",
      "tags": ["Tech", "Programming"]
  }'
  ```

- **Get All Blog Posts:**

  ```bash
  curl --location 'http://localhost:8080/posts'
  ```

- **Get a Blog Post by ID:**

  ```bash
  curl --location 'http://localhost:8080/posts/1'
  ```

- **Update a Blog Post:**

  ```bash
  curl --location --request PUT 'http://localhost:8080/posts/1' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "title": "My Updated Blog Post",
      "content": "This is the updated content of my first blog post.",
      "category": "Technology",
      "tags": ["Tech", "Programming"]
  }'
  ```

- **Delete a Blog Post:**

  ```bash
  curl --location --request DELETE 'http://localhost:8080/posts/1'
  ```

## Project Structure

```markdown
blogging-platform-api/
├── controllers/
│ └── blog_post_controller.go # HTTP handlers for blog posts
├── models/
│ └── blog_post.go # BlogPost model using GORM
├── main.go # Entry point: setup router, connect to DB, etc.
├── go.mod # Go module file
└── .env # Environment variables configuration
```

## Best Practices

- **Dependency Injection:**
  The database connection is injected into controllers, making testing and maintenance easier.

- **Error Handling:**
  Proper error checks are in place with meaningful HTTP status codes (e.g., 201, 200, 204, 400, 404).

- **Environment Configuration:**
  Use environment variables to protect sensitive configuration details (database credentials, etc.).

- **Concurrency:**
  Thanks to Go’s `net/http` package, the server handles all requests concurrently with goroutines.

- **ORM Use:**
  GORM is used to help abstract database operations and ORM benefits like auto migrations, which simplifies schema management.

## License

This project is open-source and available under the [MIT License](LICENSE).

Happy coding!
