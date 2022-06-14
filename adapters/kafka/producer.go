package kafka

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	cfg "github.com/jinagamvasubabu/golang-boilerplate/config"
	"github.com/jinagamvasubabu/golang-boilerplate/model/dto"
)

var producer Producer

type Producer interface {
	PublishMessage(ctx context.Context, message dto.Message) error
}

type syncProducer struct {
	instance sarama.SyncProducer
}

func NewSyncProducer() (p Producer, err error) {
	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{cfg.GetConfig().KafkaBrokerUrl}, config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &syncProducer{
		instance: prd,
	}, nil
}

func GetSyncProducer() (Producer, error) {
	if producer == nil {
		producer, err := NewSyncProducer()
		if err != nil {
			return producer, err
		}
	}
	return producer, nil
}

func (s syncProducer) PublishMessage(ctx context.Context, message dto.Message) error {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: cfg.GetConfig().Topic,
		Value: sarama.StringEncoder(message.Value),
	}
	p, o, err := s.instance.SendMessage(msg)
	if err != nil {
		fmt.Println("00-----000", err)
		return err
	}
	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
	return nil
}
