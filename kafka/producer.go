package kafka

import (
    "github.com/segmentio/kafka-go"
    "log"
    "context"
)

type Producer struct {
    writer *kafka.Writer
}

func NewProducer(broker, topic string) *Producer {
    return &Producer{
        writer: &kafka.Writer{
            Addr:     kafka.TCP(broker),
            Topic:    topic,
            Balancer: &kafka.LeastBytes{},
        },
    }
}

func (p *Producer) ProduceMessage(key, value []byte) error {
    err := p.writer.WriteMessages(context.Background(),
        kafka.Message{
            Key:   key,
            Value: value,
        },
    )
    if err != nil {
        log.Printf("could not write message %v", err)
        return err
    }
    return nil
}

func (p *Producer) Close() error {
    return p.writer.Close()
}
