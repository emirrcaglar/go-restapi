build:
	@go build -o bin/go-restapi cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-restapi
