server:
	go run main.go

build:
	go build -o bin.server main.go

test:
	go test