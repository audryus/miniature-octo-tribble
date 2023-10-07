build:
	@go build -o bin/octo-tribble

run: build 
	@./bin/octo-tribble

migra:  
	@go run ./migration/main.go

test:
	@go test -v ./tests/...
