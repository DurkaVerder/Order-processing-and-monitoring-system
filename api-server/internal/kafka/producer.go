// This package contains the implementation of the Kafka consumer and producer.
package kafka

import (
	"Order-processing-and-monitoring-system/common/models"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

// Producer is a wrapper around the sarama.SyncProducer to provide a more
type Producer struct {
	producer sarama.SyncProducer
	config   *sarama.Config
}

// NewProducer creates a new Producer using the given broker addresses and configuration.
func NewProducer(brokers string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{brokers}, config)
	if err != nil {
		return nil, err
	}

	return &Producer{producer, config}, nil
}

// SendMessageForCreateOrder sends a message to the Kafka topic for change status order.
func (p *Producer) SendMessageForChangeStatusOrder(topic string, order models.StatusOrder) error {
	data, err := json.Marshal(order)
	if err != nil {
		return nil
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Сообщение отправлено в топик %s, раздел %d, смещение %d\n", topic, partition, offset)
	return nil
}
