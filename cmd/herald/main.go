package main

import (
	"os"

	"github.com/cfx-cv/herald/pkg/server"
	nats "github.com/nats-io/go-nats"
)

func main() {
	conn := newNatsConn()
	server.NewServer(conn).Start()
}

func newNatsConn() *nats.Conn {
	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		panic(err)
	}

	return nc
}
