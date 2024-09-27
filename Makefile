include .env

up:
	@echo "Starting MongoDb container ..."
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stopping MongoDb container ..."
	docker-compose down

build:
	go build -o ${BINARY} ./cmd/api/

start:
	@env MONGO_DB_USERNAME=${MONGO_DB_USERNAME} MONGO_DB_PASSWORD=${MONGO_DB_PASSWORD} ./${BINARY}

restart: build start
