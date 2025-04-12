# Migration
migrate-up:
	docker compose exec api go run cmd/databases/main.go up

migrate-down:
	docker compose exec api go run cmd/databases/main.go down

migrate-status:
	docker compose exec api go run cmd/databases/main.go status

migrate-reset:
	docker compose exec api go run cmd/databases/main.go reset