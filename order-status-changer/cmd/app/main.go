package main

import (
	"order-status-changer/config"
	"order-status-changer/internal/kafka/consumer"
	"order-status-changer/internal/kafka/producer"
	"order-status-changer/internal/repository"
	"order-status-changer/internal/service"
)

func main() {
	// Initialize the configuration.
	config := config.InitConfig()

	// Initialize the repository.
	repository := repository.NewRepositoryManager(config)

	// Initialize the producer.
	producer := producer.NewProducerManager(config.Kafka.Brokers)

	// Initialize the service.
	service := service.NewServiceManager(repository, producer)

	// Initialize the consumer.
	consumer := consumer.NewConsumerManager(config.Kafka.Brokers, service)

	// Start the consumer.
	consumer.StartConsumer("orders.status")
}
