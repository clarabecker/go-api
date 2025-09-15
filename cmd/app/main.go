package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/clarabecker/estudos-go/internal/infra/akafka"
	"github.com/clarabecker/estudos-go/internal/repository"
	"github.com/clarabecker/estudos-go/internal/usecases"
	"github.com/clarabecker/estudos-go/internal/web"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	productRepository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecases.NewCreateProductUseCase(productRepository)
	listProductUseCase := usecases.NewListProductsUseCase(productRepository)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"product"}, "kafka:9092", msgChan)

	productHandler := web.NewProductHandlers(createProductUseCase, listProductUseCase)

	r := chi.NewRouter()
	r.Post("/products", productHandler.CreateProductHandler)
	r.Get("/products", productHandler.ListProductHandler)

	go func() {
		if err := http.ListenAndServe(":8000", r); err != nil {
			log.Fatalf("Erro ao iniciar HTTP server: %v", err)
		}
	}()

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
