# Makefile

# Define variables
APP_NAME = thug_geeks
APP_VERSION = 1.0
DOCKER_IMAGE = $(APP_NAME):$(APP_VERSION)

# Build the Go application
build:
	go build -o $(APP_NAME)

# Run the Go application
run: build
	./$(APP_NAME)

# Build a Docker image for the application
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Run the application in a Docker container
docker-run: docker-build
	docker run -p 8080:8080 --name $(APP_NAME) $(DOCKER_IMAGE)

# Stop and remove the Docker container
docker-stop:
	docker stop $(APP_NAME)
	docker rm $(APP_NAME)

clean:
	rm -f $(APP_NAME)
	docker rmi $(DOCKER_IMAGE)
