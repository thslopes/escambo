# Escambo POC

REST API with [Go-Chi](https://github.com/go-chi/chi) and [MongoDB Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)

This is a POC and it's not ready for production.

## Diff sample
    
```json
{
    "id": "2",
    "user_id": 10,
    "title": "article 2",
    "slug": "article-23",
    "versions": [
        {
            "article": {
                "id": "2",
                "user_id": 10,
                "title": "article 2",
                "slug": "article-21",
                "versions": null
            },
            "changes": "  db.Article{\n  \tID:       \"2\",\n  \tUserID:   10,\n  \tTitle:    \"article 2\",\n- \tSlug:     \"article-21\",\n+ \tSlug:     \"article-22\",\n  \tVersions: nil,\n  }\n",
            "user_id": 10
        },
        {
            "article": {
                "id": "2",
                "user_id": 10,
                "title": "article 2",
                "slug": "article-22",
                "versions": null
            },
            "changes": "  db.Article{\n  \tID:       \"2\",\n  \tUserID:   10,\n  \tTitle:    \"article 2\",\n- \tSlug:     \"article-22\",\n+ \tSlug:     \"article-23\",\n  \tVersions: nil,\n  }\n",
            "user_id": 10
        }
    ],
    "elapsed": 10
}
```  

## MongoDB Docker
```sh
docker run -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=test -e MONGO_INITDB_ROOT_PASSWORD=123 -d mongo:latest
```

## Load DB

```sh
go run cmd/mock/main.go
```

## Running the app

```sh
go run cmd/main.go
```

## Make

```sh
make mongo-docker
make load-db
make run
```

## Endpoints

### /articles

#### GET

```sh
curl --location --request GET 'http://localhost:3333/articles'
```

#### POST

```sh
curl "http://localhost:3333/articles" \
-d '
{
    "title": "Article 1",
    "slug": "article-1",
    "user_id": 10
}'
```

### /articles/{id}

#### GET

```sh
curl --location --request GET 'http://localhost:3333/articles/2'
```

#### PUT

```sh
curl --location --request PUT 'http://localhost:3333/articles/2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Article 2",
    "slug": "article-2",
    "user_id": 10
}'
```