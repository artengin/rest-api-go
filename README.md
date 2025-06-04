# REST API in Go for Managing Person Model

## Description

This project is an example of a RESTful CRUD service implemented in Go, inspired by Clean Architecture principles.  
The service allows you to create, retrieve, update, and delete records about people (Person) in a PostgreSQL database.

## Tech Stack

- Language: Go
- HTTP Framework: [Echo](https://echo.labstack.com/)
- Database Access: [gocraft/dbr](https://github.com/gocraft/dbr)
- Logging: [logrus](https://github.com/sirupsen/logrus)
- Validation: [go-playground/validator](https://github.com/go-playground/validator)
- Database: PostgreSQL

## Architecture

The project is divided into 4 layers:

- **internal/app** — model definition (Person)
- **internal/http** — HTTP handlers
- **internal/logic** — business logic (use cases)
- **internal/repository/postgres** — data access layer (PostgreSQL)

## Person Model

```go
type Person struct {
    ID        int64     `json:"id"`
    Email     string    `json:"email"`
    Phone     string    `json:"phone"`
    FirstName string    `json:"firstName"`
    LastName  string    `json:"lastName"`
    UpdatedAt time.Time `json:"updatedAt"`
    CreatedAt time.Time `json:"createdAt"`
}
```

## Running the Project

1. **Clone the repository:**

2. **Create a `.env` file** in the project root 

3. **Start PostgreSQL and create the database.**

4. **Run the application:**
   ```bash
   go run .
   ```

   On the first run, the migration for the `person` table will be applied automatically.

## API

### Get list of Persons

```
GET /person
```

or

```
GET /person?limit=10&offset=0&search=John
```

### Get a single Person

```
GET /person/{id}
```

### Create a Person

```
POST /person
Content-Type: application/json

{
  "email": "test@example.com",
  "phone": "+79991234567",
  "firstName": "John",
  "lastName": "Doe"
}
```

### Update a Person

```
PUT /person/{id}
Content-Type: application/json

{
  "email": "test2@example.com",
  "phone": "+79991234568",
  "firstName": "Peter",
  "lastName": "Smith"
}
```

### Delete a Person

```
DELETE /person/{id}
```