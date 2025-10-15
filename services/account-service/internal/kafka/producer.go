package kafka

import "github.com/IBM/sarama"

func NewProducer(brokers string) (sarama.SyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	return sarama.NewSyncProducer([]string{brokers}, cfg)
}
