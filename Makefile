.PHONY: run build test clean fmt vet lint

run:
	@go run main.go

build:
	@go build -o godiscopwn main.go

test:
	@go test ./...

clean:
	@go clean
	@rm -f godiscopwn

fmt:
	@go fmt ./...

vet:
	@go vet ./...

lint:
	@golangci-lint run
