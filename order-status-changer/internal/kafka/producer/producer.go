// This package contains the Producer interface and ProducerManager struct, which are used to send messages to Kafka topics.
package producer

import (
	"Order-processing-and-monitoring-system/common/models"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

// Producer is a wrapper around the sarama.SyncProducer to provide a more
type Producer interface {
	SendMessageForAnalytics(topic string, report models.Report) error
}

// ProducerManager is a wrapper around the sarama.SyncProducer to provide a more
type ProducerManager struct {
	producer sarama.SyncProducer
	config   *sarama.Config
}

// NewProducerManager creates a new ProducerManager using the given broker addresses.
func NewProducerManager(brokers string) *ProducerManager {
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

// SendMessageForAddOrder sends a message to the Kafka topic for adding an order.
func (p *ProducerManager) SendMessageForAnalytics(topic string, report models.Report) error {
	data, err := json.Marshal(report)
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
