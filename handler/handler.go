package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "test_from_Messaggio/repository"
    "test_from_Messaggio/kafka"
    "test_from_Messaggio/model"
)

type Handler struct {
    repo          *repository.Repository
    kafkaProducer *kafka.Producer
}

func NewHandler(repo *repository.Repository, kafkaProducer *kafka.Producer) *Handler {
    return &Handler{repo: repo, kafkaProducer: kafkaProducer}
}

func (h *Handler) CreateMessage(c *gin.Context) {
    var message struct {
        Content string `json:"content"`
    }
    if err := c.ShouldBindJSON(&message); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Create and save the message
    msg := &model.Message{
        Content:   message.Content,
        Processed: false,
    }
    if err := h.repo.SaveMessage(msg); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save message"})
        return
    }

    // Produce message to Kafka
    if err := h.kafkaProducer.ProduceMessage([]byte("key"), []byte(message.Content)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not produce message"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Message created"})
}

func (h *Handler) GetProcessedMessages(c *gin.Context) {
    messages, err := h.repo.GetProcessedMessages()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve messages"})
        return
    }
    c.JSON(http.StatusOK, messages)
}
