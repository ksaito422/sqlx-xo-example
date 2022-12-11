run:
	go run ./...
mod:
	go mod tidy
schema:
	xo schema mysql://admin:secret@127.0.0.1:3306/sqlx_xo -o src/models
