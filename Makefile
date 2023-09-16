DOCKER_IMAGE_NAME := my_docker_image
DOCKERFILE := ./Dockerfile
PATH_TO_PUSH := my_path

build:
	docker build -t $(DOCKER_IMAGE_NAME) -f $(DOCKERFILE) .

push:
	docker push $(PATH_TO_PUSH)