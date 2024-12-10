package kafka

import (
	"encoding/json"

	"github.com/IBM/sarama" // Updated import path
)

type Producer struct {
    producer sarama.SyncProducer
}

type Message struct {
    SenderID string `json:"sender_id"`
    Content  string `json:"content"`
	ChatRoomId string `json:"chat_room_id"`
}

func NewProducer(brokers []string) (*Producer, error) {
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5

    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        return nil, err
    }

    return &Producer{producer: producer}, nil
}

func (p *Producer) SendMessage(topic string, message *Message) error {
    msgBytes, err := json.Marshal(message)
    if err != nil {
        return err
    }

    msg := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.StringEncoder(msgBytes),
    }

    _, _, err = p.producer.SendMessage(msg)
    return err
}

func (p *Producer) Close() error {
    return p.producer.Close()
}