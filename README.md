# Escambo POC

REST API with [Go-Chi](https://github.com/go-chi/chi) and [MongoDB Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)

This is a POC and it's not ready for production.

## MongoDB Docker
```sh
docker run -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=test -e MONGO_INITDB_ROOT_PASSWORD=123 -d mongo:latest
```

## Running the app

```sh
go run cmd/main.go
```

## Make

```sh
make mongo-docker
make run
```