package handler

import (
    "github.com/gin-gonic/gin"
    "test_from_Messaggio/model"
    "test_from_Messaggio/repository"
    "test_from_Messaggio/kafka"
    "net/http"
)

type Handler struct {
    Repo   repository.Repository
    KafkaProducer kafka.Producer
}

func NewHandler(repo repository.Repository, kafkaProducer kafka.Producer) *Handler {
    return &Handler{Repo: repo, KafkaProducer: kafkaProducer}
}

func (h *Handler) CreateMessage(c *gin.Context) {
    var message model.Message
    if err := c.ShouldBindJSON(&message); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.Repo.SaveMessage(&message); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err := h.KafkaProducer.SendMessage(&message); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Message created successfully"})
}

func (h *Handler) GetProcessedMessages(c *gin.Context) {
    messages, err := h.Repo.GetProcessedMessages()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, messages)
}
