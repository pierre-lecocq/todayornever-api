BIN_NAME=todayornever-api

.PHONY: build test coverage run start clean

all: build

test:
	go test ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

build:
	go build -o ./$(BIN_NAME) .

run:
	go run main.go

start:
	./$(BIN_NAME) # env $(cat .env | xargs) ./$(BIN_NAME)

clean:
	rm -f ./$(BIN_NAME) coverage.out

# Database migrations

db-up:
	migrate -path ./app/migrations -database sqlite3://./$(BIN_NAME).db up

db-down:
	migrate -path ./app/migrations -database sqlite3://./$(BIN_NAME).db down

db-backup:
	sqlite3 ./$(BIN_NAME).db ".backup '$(BIN_NAME).backup.db'"

# Docker build

docker-build:
	docker build -t $(BIN_NAME):latest .

docker-run:
	docker run \
		--port 8080:8080 \
		--mount type=bind,source="$(pwd)/data",target=/app
		$(BIN_NAME):latest

docker-clean:
	docker rmi -f $(BIN_NAME):latest
