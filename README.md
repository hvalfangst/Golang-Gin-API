# Golang Gin API with Bun ORM




## Requirements

* x86-64
* Linux/Unix
* [Golang](https://go.dev/)
* [Docker](https://www.docker.com/products/docker-desktop/)

## Startup

The script "up" creates our DB container, compiles and executes our binary:
```
1. docker-compose -f db/cars/docker-compose.yml up -d
2. go build -o gin_api src/main.go
3. ./gin_api
```

## Shutdown

The script "down" removed our DB container:
```
1. docker-compose -f db/cars/docker-compose.yml down
```

## Postman Collection

The repository includes a Postman collection in the 'postman' directory.


