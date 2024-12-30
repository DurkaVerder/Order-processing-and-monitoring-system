// This package contains the implementation of the Kafka consumer and producer.
package kafka

type Kafka struct {
	Producer Producer
}

// NewKafka creates a new Kafka
func NewKafka(producer Producer) *Kafka {
	return &Kafka{producer}
}
