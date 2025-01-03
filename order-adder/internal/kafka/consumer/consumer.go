package consumer

import (
	"encoding/json"
	"log"
	"order-adder/internal/kafka"
	"order-adder/internal/models"
	"order-adder/internal/service"
	"time"

	"github.com/IBM/sarama"
)

// Consumer is a wrapper around the sarama.Consumer to provide a more
type Consumer interface {
	StartConsumer(topic string)
}

// ConsumerManager is a wrapper around the sarama.Consumer to provide a more
type ConsumerManager struct {
	consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer
	config            *sarama.Config
	service           service.Service
}

// NewConsumerManager creates a new ConsumerManager using the given broker addresses.
func NewConsumerManager(brokers string, service service.Service) *ConsumerManager {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	for i := 0; i < kafka.MaxRetries; i++ {

		consumer, err := sarama.NewConsumer([]string{brokers}, config)
		if err != nil {
			log.Println("Successful create consumer")
			return &ConsumerManager{consumer: consumer, config: config, service: service}
		}
		log.Println("Error creating consumer: ", err)
		time.Sleep(5 * time.Second)
	}
	log.Fatal("Error creating consumer")
	return nil
}

// SubscribeTopic subscribes to the given topic.
func (c *ConsumerManager) subscribeTopic(topic string) error {
	var err error
	c.partitionConsumer, err = c.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}
	return nil
}

// StartConsumer starts the consumer.
func (c *ConsumerManager) StartConsumer(topic string) {
	if err := c.subscribeTopic(topic); err != nil {
		log.Fatal("Error creating partition consumer: ", err)
	}
	defer c.partitionConsumer.Close()

	for {

		select {
		case msg := <-c.partitionConsumer.Messages():

			log.Printf("Message: %s\n", string(msg.Value))

			order := models.Order{}
			if err := json.Unmarshal(msg.Value, &order); err != nil {
				log.Printf("Error unmarshalling the message: %s\n", err.Error())
			}

			if err := c.service.AddOrder(order); err != nil {
				log.Printf("Error adding the order: %s\n", err.Error())
			} else {
				log.Printf("Order added successfully\n")
			}

		case err := <-c.partitionConsumer.Errors():
			log.Printf("Error: %s\n", err.Error())
		}
	}
}
