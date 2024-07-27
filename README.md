This repository contains a simple CRUD (Create, Read, Update, Delete) application built with Golang, using Gorilla Mux for routing and Go SQL Driver. The application manages user data, providing endpoints to perform CRUD operations.

## Dependency
- [Gorilla Mux 🦍](https://github.com/gorilla/mux)
- [Go SQL Driver](https://github.com/go-sql-driver/mysql)

## Features
- Create a new user
- Read user data
- Update existing user data
- Delete a user

## Installation
Clone the repository:

```bash
git clone https://github.com/yourusername/simple-golang-crud.git
cd golang-crud-user
```

Install dependencies:
```bash
go mod tidy
```

## Usage
Run the application:
```bash
go run main.go
```

## Testing
API Endpoints:

Create a user:
```
curl -X POST -H "Content-Type: application/json" -d '{"Name":"John Doe","Email":"john@example.com"}' http://localhost:8090/user
```

Read a user
```
curl http://localhost:8090/user/{id}
```

Update user
```
curl -X PUT -H "Content-Type: application/json" -d '{"Name":"Jane Doe","Email":"jane@example.com"}' http://localhost:8090/user/{id}
```

Delete user
```
curl -X DELETE http://localhost:8090/user/{id}
```


## License
This project is licensed under the MIT License.
