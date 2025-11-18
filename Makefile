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
	docker build -t $(APP_NAME)-image .
	docker run -d -p 8081:8080 --name $(APP_NAME) $(APP_NAME)-image

.PHONY: docker-stop
docker-stop:
	docker stop $(APP_NAME) || true
	docker rm $(APP_NAME) || true