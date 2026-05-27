# Future Message

A Go backend service for a scheduled letter/time-capsule application.

This repository provides an API for user registration, login, and logout using PostgreSQL for persistence and Redis for session management. The project includes the core auth flow and a future-ready data model for scheduled letters.

## What it does

- `POST /auth/register` — register a new user
- `POST /auth/login` — login with email and password
- `POST /auth/logout` — destroy the current session
- stores users in PostgreSQL
- uses Redis via [alexedwards/scs](https://github.com/alexedwards/scs) for secure session cookies

## Architecture

- `cmd/api` — HTTP server entry point
- `internal/api` — router, handlers and middleware
- `internal/service` — business logic
- `internal/repository` — database access
- `internal/models` — domain entities (`User`, `Letter`)
- `pkg/database` — PostgreSQL connection helper
- `pkg/session` — Redis session manager
- `internal/db/migrations` — database schema migrations

## Requirements

- Go 1.26
- PostgreSQL
- Redis

## Local development

Start PostgreSQL and Redis with Docker Compose:

```bash
make docker-up
```

Set your environment variables in a `.env` file or export them directly.

Example `.env`:

```env
DB_URL=postgres://postgres:123456@localhost:5433/time_capsule?sslmode=disable
PORT=8080
```

Run the API server:

```bash
make run-api
```

The server listens on `:8080` by default.

## Database migrations

Migrations are stored in `internal/db/migrations`.

Run the example migration command:

```bash
make migrate-up
```

> The `make migrate-up` target uses the local example Postgres URL. Update it if your database settings differ.

## API Endpoints

### Register

`POST /auth/register`

Request body:

```json
{
  "email": "user@example.com",
  "username": "johndoe",
  "password": "secret"
}
```

### Login

`POST /auth/login`

Request body:

```json
{
  "email": "user@example.com",
  "password": "secret"
}
```

### Logout

`POST /auth/logout`

No body required.

## Environment variables

- `DB_URL` — PostgreSQL connection string
- `PORT` — optional HTTP port (defaults to `8080`)

## Notes

- The current implementation covers authentication and session management.
- The `Letter` model and migration are present, and protected letter endpoints are scaffolded in `internal/api/router.go`.
- `cmd/worker` exists as a worker command placeholder for future background processing.

## Useful commands

```bash
make docker-up   # start Postgres and Redis
make run-api     # run the API server
make test        # run Go tests
make dev         # run hot reload with air (if installed)
```

## Project status

This is an early-stage backend service for a future-letter app. Authentication is implemented, while scheduled letters and worker processing are planned and partially scaffolded.

## Future work

- implement protected `letters` endpoints for creating, listing, and retrieving scheduled letters
- wire the `Letter` model into the API and repository layer
- add a worker process to deliver letters at the scheduled `deliver_at` time
- add end-to-end tests for auth, letter creation, and delivery workflows
