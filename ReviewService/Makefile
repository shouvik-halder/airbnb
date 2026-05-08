DB_MIGRATION_FOLDER = db/migrations
run:
	go run main.go

migrate-create:
	goose -dir $(DB_MIGRATION_FOLDER) create $(name) sql

migrate-up: # gmake migrate-up -url="db_connection_url"
	goose -dir $(DB_MIGRATION_FOLDER) mysql "$(url)" up
