DOCKER_IMAGE_NAME := my_docker_image
DOCKERFILE := ./Dockerfile

build:
	docker build -t $(DOCKER_IMAGE_NAME) -f $(DOCKERFILE) .