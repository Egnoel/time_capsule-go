run-api:
	go run ./cmd/api

run-worker:
	go run ./cmd/worker

migrate-up:
	migrate -path internal/db/migrations -database "postgres://postgres:123456@localhost:5433/time_capsule?sslmode=disable" up

test:
	go test ./... -v

docker-up:
	docker compose up -d

dev:
	air