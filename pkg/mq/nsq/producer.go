package nsq

import (
	"github.com/nsqio/go-nsq"
)

type Producer struct {
	*nsq.Producer
}

func NewProducer(addr string) (*Producer, error) {
	producer, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	p := new(Producer)
	p.Producer = producer
	return p, nil
}
