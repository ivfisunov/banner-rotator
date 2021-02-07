lint:
	golangci-lint run ./...

test:
	go test -race -count 100 ./...

build:
	go build  -v -o ./bin/rotator ./cmd/rotator

run:
	docker-compose up

goose-up:
	goose -dir=migrations postgres "user=banner password=123456 dbname=banner sslmode=disable" up

goose-down:
	goose -dir=migrations postgres "user=banner password=123456 dbname=banner sslmode=disable" down
