mongo-docker:
	docker run -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=test -e MONGO_INITDB_ROOT_PASSWORD=123 -d mongo:latest

run:
	go run cmd/main.go