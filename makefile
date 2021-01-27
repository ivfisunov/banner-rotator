goose-up:
	goose -dir=migrations postgres "user=banner password=123456 dbname=banner sslmode=disable" up

goose-down:
	goose -dir=migrations postgres "user=banner password=123456 dbname=banner sslmode=disable" down
