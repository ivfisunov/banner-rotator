lint:
	golangci-lint run ./...

test:
	go test ./...

build:
	go build  -v -o ./bin/rotator ./cmd/rotator

run:
	go run cmd/rotator/main.go

goose-up:
	goose -dir=migrations postgres "user=banner password=123456 dbname=banner sslmode=disable" up

goose-down:
	goose -dir=migrations postgres "user=banner password=123456 dbname=banner sslmode=disable" down
