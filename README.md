# Escambo POC

## MongoDB Docker
```sh
docker run -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=test -e MONGO_INITDB_ROOT_PASSWORD=123 -d mongo:latest
```

## Running the app

```sh
go run cmd/main.go
```