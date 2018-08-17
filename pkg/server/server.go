package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	nats "github.com/nats-io/go-nats"

	"github.com/cfx-cv/herald/pkg/herald"
	hnats "github.com/cfx-cv/herald/pkg/nats"
)

const (
	publishURL string = "/publish"
)

type Server struct {
	broker herald.Broker
}

func NewServer(conn *nats.Conn) *Server {
	broker := hnats.NewBroker(conn)
	return &Server{broker: broker}
}

func (s *Server) Start() {
	router := mux.NewRouter()
	router.HandleFunc(publishURL, s.publish).Methods("POST")

	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal(err)
	}
}
