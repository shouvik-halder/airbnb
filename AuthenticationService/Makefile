run:
	go run main.go

migrate-up:
	goose -dir db/migrations mysql "${DB_URL}" up

build:
	go build -o app