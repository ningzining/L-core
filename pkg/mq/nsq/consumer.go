package nsq

import (
	"github.com/nsqio/go-nsq"
)

type callbackFunc func(msg *nsq.Message) error

type Consumer struct {
	*nsq.Consumer
}

func NewConsumer(topic, channel, addr string, callback callbackFunc) (*Consumer, error) {
	consumer, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		return callback(msg)
	}))
	if err := consumer.ConnectToNSQD(addr); err != nil {
		return nil, err
	}
	c := new(Consumer)
	c.Consumer = consumer
	return c, nil
}
