// This package contains the implementation of the Kafka consumer and producer.
package kafka

import (
	"Order-processing-and-monitoring-system/common/models"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

// Producer is a generic interface for producing messages to Kafka.
type Producer interface {
	SendMessageForChangeStatusOrder(topic string, order models.StatusOrder) error
	SendMessageForCreateOrder(topic string, order models.Order) error
}

// ProducerManager is a Kafka producer that sends messages to a Kafka topic.
type ProducerManager struct {
	producer sarama.SyncProducer
	config   *sarama.Config
}

// NewProducer creates a new Producer using the given broker addresses and configuration.
func NewProducer(brokers string) *ProducerManager {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{brokers}, config)
	if err != nil {
		log.Fatal("Error creating producer: ", err)
	}

	return &ProducerManager{producer, config}
}

// SendMessageForCreateOrder sends a message to the Kafka topic for change status order.
func (p *ProducerManager) SendMessageForChangeStatusOrder(topic string, order models.StatusOrder) error {
	data, err := json.Marshal(order)
	if err != nil {
		return nil
	}

	if err = p.sendMessage(topic, data); err != nil {
		return err
	}
	return nil
}

// SendMessageForCreateOrder sends a message to the Kafka topic for create order.
func (p *ProducerManager) SendMessageForCreateOrder(topic string, order models.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return nil
	}

	if err = p.sendMessage(topic, data); err != nil {
		return err
	}
	return nil
}

// sendMessage sends a message to the Kafka topic.
func (p *ProducerManager) sendMessage(topic string, data []byte) error {
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
