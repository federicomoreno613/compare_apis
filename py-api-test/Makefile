# Variables
IMAGE_NAME := api-test
CONTAINER_NAME := api-test-app
NETWORK_NAME := api-test-app-net
TEST_COMMAND := pytest

.PHONY: docker-build
docker-build:
	docker build -t $(IMAGE_NAME) .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8081 --name $(CONTAINER_NAME) --network $(NETWORK_NAME) $(IMAGE_NAME)

.PHONY: test
test:
	docker run -v $(PWD):/app --network $(NETWORK_NAME) $(IMAGE_NAME) $(TEST_COMMAND)

.PHONY: clean
clean:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)
