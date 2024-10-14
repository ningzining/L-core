package nsq

import (
	"github.com/nsqio/go-nsq"
)

func NewProducer(addr string) (*nsq.Producer, error) {
	return nsq.NewProducer(addr, nsq.NewConfig())
}
