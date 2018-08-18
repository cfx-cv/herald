package nats

import (
	"time"

	nats "github.com/nats-io/go-nats"
)

type Broker struct {
	conn *nats.Conn
}

func NewBroker(conn *nats.Conn) *Broker {
	return &Broker{conn: conn}
}

func (nb *Broker) Publish(topic, message string) error {
	return nb.conn.Publish(topic, []byte(message))
}

func (n *Broker) Request(topic, message string) (interface{}, error) {
	return n.conn.Request(topic, []byte(message), 1*time.Second)
}
