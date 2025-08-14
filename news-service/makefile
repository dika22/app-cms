APP=news-service
APP_EXECUTABLE=${APP}

serve-http:
	go run main.go serve-http

migrate:
	go run main.go migrate-db

test:
	go test -v ./...

swag-init:
	swag init
