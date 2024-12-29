// // This package contains the implementation of the Kafka consumer and producer.
package kafka

// import "github.com/IBM/sarama"

// // Consumer is a wrapper around the sarama.Consumer to provide a more
// type Consumer struct {
// 	consumer sarama.Consumer
// 	config   *sarama.Config
// }

// // NewConsumer creates a new Consumer using the given broker addresses and configuration.
// func NewConsumer(brokers string) (*Consumer, error) {
// 	config := sarama.NewConfig()
// 	config.Consumer.Return.Errors = true

// 	consumer, err := sarama.NewConsumer([]string{brokers}, config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Consumer{consumer, config}, nil
// }
