# Variables
APP_NAME=packman-app

# Default target
help:
	@echo "Available targets:"
	@echo "  make run         		  - Run the Go application"
	@echo "  make docker-run          - Run the Go application on Docker container"
	@echo "  make docker-stop         - Stop the Go application Docker container"

.PHONY: run
run:
	go run main.go

.PHONY: docker-run
docker-run:
	docker build -t $(APP_NAME) .
	docker network inspect packman-network >/dev/null 2>&1 || docker network create packman-network
	docker run --network packman-network -p 8080:8080 --name $(APP_NAME) -e PACKMAN_API_BASE_URL=http://localhost:8081 $(APP_NAME)

.PHONY: docker-stop
docker-stop:
	docker stop $(APP_NAME) || true
	docker rm $(APP_NAME) || true