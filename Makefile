build:
	@go build -o bin/byteknot cmd/byteknot/main.go

test:
	@go test -v ./...

run: build
	@./bin/byteknot