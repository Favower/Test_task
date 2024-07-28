package kafka

import (
    "context"
    "github.com/segmentio/kafka-go"
    "test_from_Messaggio/model"
    "encoding/json"
)

type Producer struct {
    Writer *kafka.Writer
}

func NewProducer(broker, topic string) *Producer {
    return &Producer{
        Writer: &kafka.Writer{
            Addr:     kafka.TCP(broker),
            Topic:    topic,
            Balancer: &kafka.LeastBytes{},
        },
    }
}

func (p *Producer) SendMessage(message *model.Message) error {
    msgBytes, err := json.Marshal(message)
    if err != nil {
        return err
    }

    return p.Writer.WriteMessages(context.Background(),
        kafka.Message{
            Key:   []byte(string(message.ID)),
            Value: msgBytes,
        },
    )
}
