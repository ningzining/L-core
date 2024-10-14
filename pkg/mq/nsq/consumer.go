package nsq

import (
	"github.com/nsqio/go-nsq"
)

type ConsumerHandler struct {
}

func NewConsumerHandler() nsq.Handler {
	return &ConsumerHandler{}
}

func (d *ConsumerHandler) HandleMessage(message *nsq.Message) error {
	return nil
}

func NewConsumer(topic, channel, addr string, handler nsq.Handler) (*nsq.Consumer, error) {
	consumer, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	consumer.AddHandler(handler)
	if err := consumer.ConnectToNSQD(addr); err != nil {
		return nil, err
	}

	return consumer, nil
}
