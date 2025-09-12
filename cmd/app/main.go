package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/clarabecker/estudos-go/internal/infra/akafka"
	"github.com/clarabecker/estudos-go/internal/repository"
	"github.com/clarabecker/estudos-go/internal/usecases"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3600)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	productRepository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecases.NewCreateProductUseCase(productRepository)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		var dto usecases.CreateProductInputDTO
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			log.Printf("Erro ao fazer unmarshal da mensagem: %v", err)
			continue
		}

		_, err = createProductUseCase.Execute(dto)
		if err != nil {
			log.Printf("Erro ao executar o use case: %v", err)
		}
	}
}
