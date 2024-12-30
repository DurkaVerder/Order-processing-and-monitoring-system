package main

import (
	"order-adder/config"
	"order-adder/internal/kafka/consumer"
	"order-adder/internal/kafka/producer"
	"order-adder/internal/repository"
	"order-adder/internal/service"
)

func main() {
	// Initialize the configuration and repository.
	config := config.InitConfig()

	// Initialize the repository.
	repository := repository.NewRepositoryManager(config)

	// Initialize the Kafka producer.
	producer := producer.NewProducerManager(config.Kafka.Brokers)

	// Initialize the service.
	service := service.NewServiceManager(repository, producer)

	// Initialize the Kafka consumer.
	consumer := consumer.NewConsumerManager(config.Kafka.Brokers, service)

	consumer.StartConsumer("orders.new")
}
