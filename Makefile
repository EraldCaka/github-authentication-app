include .env
export

build:
	@go build -o bin/api cmd/main.go

run: build
	@./bin/api

migrate:
	@migrate -path ./db/migrations -database "$(DB_LINK)" -verbose up

drop:
	@migrate -path db/migrations -database "$(DB_LINK)" -verbose down

create-migration-sql:
	@migrate create -ext sql -dir db/migrations github_user_table_and_commit_history_table