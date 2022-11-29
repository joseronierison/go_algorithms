install:
	go mod tidy
	go mod download

test:
	go test -v ./...