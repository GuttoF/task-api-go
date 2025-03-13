.PHONY: build run test clean docker-build docker-run

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

test:
	go test -v ./...

clean:
	rm -rf bin/

docker-build:
	docker build -t go-todo-api .

docker-run:
	docker run -p 8080:8080 go-todo-api
