include .env
export

build:
	@go build -o bin/api cmd/main.go

run:
	@./bin/api

migrate:
	@migrate -path db/migrations -database "$(DB_URL)" -verbose up

drop:
	@migrate -path db/migrations -database "$(DB_URL)" -verbose down

create-migration-sql:
	@migrate create -ext sql -dir db/migrations github_user_table_and_commit_history_table

migrate-force-version:
	@migrate -path ./db/migrations -database "$(DB_LINK)" -verbose force 20240224023216

tidy:
	go mod tidy

worker:
	@go build -o bin/worker cmd/worker.go
	@./bin/worker

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down