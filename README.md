# Go API with MySQL and Kafka

A simple Go API for managing products, integrated with MySQL and Kafka, containerized using Docker and Docker Compose.

## Technologies

- Go
- MySQL 5.7
- Kafka + Zookeeper
- Docker & Docker Compose
- Chi router

## Prerequisites

Make sure you have installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- (Optional) [Go](https://golang.org/) if you want to run locally


## Setup 

1. Clone the repository:

```bash
git clone https://github.com/clarabecker/go-api.git
cd go-api
```

2. Build and start the containers:
```bash
docker compose up -d --build
```
## Run the Application

1. Enter Go container
```bash
docker compose exec goapp bash
```

2. Run API
```bash
go run cmd/app/main.go
```
## API Endpoints

* POST /products: Create a new product (sends event to Kafka).

* GET /products: List all products.

