APP_NAME := api-test
DOCKER_IMAGE := federicomoreno613/$(APP_NAME):latest

.PHONY: all clean build docker run stop

all: clean build docker run

clean:
	@echo "Cleaning up"
	@docker stop $(APP_NAME) >/dev/null 2>&1 || true
	@docker rm $(APP_NAME) >/dev/null 2>&1 || true
	@docker rmi $(DOCKER_IMAGE) >/dev/null 2>&1 || true

build:
	@echo "Building binary"
	@go build -o $(APP_NAME) main.go

docker:
	@echo "Building Docker image"
	@docker build -t $(DOCKER_IMAGE) .

run:
	@echo "Running service"
	@docker run -d --name $(APP_NAME) -p 8080:8080 $(DOCKER_IMAGE)

stop:
	@echo "Stopping service"
	@docker stop $(APP_NAME)
	@docker rm $(APP_NAME)

test:
	@echo "Running tests"
	@go test -v ./...