include .env

up:
	@echo "Starting MongoDb container ..."
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stoping MongoDb container ..."
	docker-compose down

build:
	go build -o ${BINARY} ./cmd/api/

start:
	./${BINARY}

restart: build start
