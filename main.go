package main

import (
    "context"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v4"
    "app/config"
    "app/handler"
    "app/repository"
    "app/kafka"
)

func main() {
    cfg, err := config.LoadConfig()
    if (err != nil) {
        log.Fatalf("could not load config: %v", err)
    }

    db, err := pgx.Connect(context.Background(), cfg.PostgresURL)
    if err != nil {
        log.Fatalf("could not connect to postgres: %v", err)
    }
    defer db.Close(context.Background())

    repo := repository.NewRepository(db)
    kafkaProducer := kafka.NewProducer(cfg.KafkaBroker, cfg.KafkaTopic)

    h := handler.NewHandler(repo, kafkaProducer)

    r := gin.Default()
    r.POST("/messages", h.CreateMessage)
    r.GET("/messages/processed", h.GetProcessedMessages)

    if err := r.Run(":8080"); err != nil {
        log.Fatalf("could not run server: %v", err)
    }
}
