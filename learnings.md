# Learnings

## Why Use Gorilla Mux Instead of Gin?

**Gorilla Mux** is a lightweight and flexible router built on top of Go’s standard `net/http` library. Here are some reasons for choosing it:

- **Simplicity & Control:** It allows you full control over the request handling without extra abstraction or opinionated behavior. You can easily design your middleware, logging, or error handling.
- **Flexibility:** It does one thing—routing—very well, and integrates seamlessly with other libraries (like GORM for database interactions).
- **Standard Concurrency:** Go's built-in HTTP server spawns a goroutine for every request, making your API asynchronous and capable of handling many requests concurrently by default.

While **Gin** provides high performance along with many built-in features such as JSON binding and middleware support, for a simple CRUD API where explicit control and clarity are valued, Gorilla Mux offers a minimal yet robust solution.

## Go Networking & Concurrency

- **Concurrent by Default:**
  The standard `net/http` package automatically spawns a new goroutine per incoming request, so all servers built using Gorilla Mux (or Gin) are asynchronous at their core—making them inherently concurrent.

- **Scalability:**
  Go’s lightweight goroutines and efficient scheduler provide excellent scalability. For most web applications, the concurrent model offered by Go's built-in libraries is more than sufficient.

- **Other Frameworks:**
  Other popular frameworks include:

  - **Gin:** A feature-rich and high-performance framework that comes with built-in JSON binding, middleware, and more.
  - **Echo & Fiber:** Echo is similar to Gin in terms of features, while Fiber (built on the fasthttp package) is designed for ultra-low latency.

  The “best” framework depends on project needs. For this project, Gorilla Mux paired with the standard `net/http` and GORM covers all requirements while remaining simple and modular.
