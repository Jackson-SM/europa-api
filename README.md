# Europa

REST API built with Go for user management. Uses Gin, GORM, and PostgreSQL.

## Stack

- **Go 1.24** — Gin for routing, GORM for ORM
- **PostgreSQL** — via Docker
- **bcrypt** — password hashing
- **UUID** — primary keys

## Project Structure

```
cmd/api/          → entrypoint
internal/
├── config/       → env config loader
├── controllers/  → request handlers
├── database/     → connection + migrations
├── domain/       → models
├── dto/          → request/response schemas
├── repositories/ → data access layer
└── routes/       → route definitions
```

## Getting Started

### Prerequisites

- Go 1.24+
- Docker & Docker Compose

### Setup

1. Clone the repo

```bash
git clone https://github.com/Jackson-SM/Europa.git
cd Europa
```

2. Create a `.env` file

```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=europa
DATABASE_URL=postgres://postgres:postgres@localhost:5432/europa?sslmode=disable
PORT=5432
```

3. Start the database

```bash
docker compose up db -d
```

4. Run the API

```bash
go run cmd/api/main.go
```

Or run everything with Docker:

```bash
docker compose up --build
```

The API starts on `http://localhost:3030`.

## API Endpoints

Base path: `/api/v1`

| Method | Route         | Description      |
|--------|---------------|------------------|
| POST   | `/users`      | Create a user    |
| GET    | `/users/:id`  | Get user by ID   |

### Create User

```bash
curl -X POST http://localhost:3030/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John", "email": "john@mail.com", "password": "12345678"}'
```

### Get User

```bash
curl http://localhost:3030/api/v1/users/<uuid>
```

## License

MIT — see [LICENSE](LICENSE)
