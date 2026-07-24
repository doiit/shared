package events

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid/v5"

	"github.com/IBM/sarama"
)

type EventProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewEventProducer(brokers []string, topic string) (*EventProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return &EventProducer{
		producer: producer,
		topic:    topic,
	}, nil
}

func (p *EventProducer) PublishEvent(userID uuid.UUID, eventType string, timestamp time.Time, event any) error {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(userID.String()),
		Value: sarama.ByteEncoder(eventBytes),
		Headers: []sarama.RecordHeader{
			{Key: []byte("event_type"), Value: []byte(eventType)},
			{Key: []byte("timestamp"), Value: []byte(timestamp.String())},
		},
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	log.Printf("Event published to partition %d at offset %d", partition, offset)
	return nil
}

// Close producer gracefully
func (p *EventProducer) Close() error {
	return p.producer.Close()
}
