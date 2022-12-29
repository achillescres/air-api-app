include .env

PROJECT_NAME=$(shell basename "$(PWD)")
DOCKER_USERNAME=achillescres

export PROJECT_ABS_PATH := /home/achillescres/GolandProjects/air-api-app

# Utils
clean:
	go clean
	rm -f saina-api

i:
	go mod tidy
	go mod download

vet:
	go vet


# Build
build:
	GOOS=linux go build -o saina-api cmd/main.go

run:
	./saina-api

# Naked(without docker)
bnr: build run

SAINA_TRASH_STAGE=saina_trash_stage

# Docker
docker: clean
	# Check and prepare module
	go mod tidy

	# Remove previous images
	docker image rm ${DOCKER_USERNAME}/saina-api || true

	# Build image
	docker build --tag ${DOCKER_USERNAME}/saina-api --build-arg TRASH_SIGN=$(SAINA_TRASH_STAGE) .

	# Run image
	docker run --rm --network=host -p 7771:7771 --name saina-api ${DOCKER_USERNAME}/saina-api

# docker rmi $(docker images -a -q --filter "label=sign=SAINA_TRASH_STAGE")

mup:
	migrate -path ./migrations -database 'postgres://postgres:Cerfvcsa@localhost:5432/dev' up

mdown:
	migrate -path ./migrations -database 'postgres://postgres:Cerfvcsa@localhost:5432/dev' up