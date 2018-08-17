package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cfx-cv/herald/pkg/herald"
)

func (s *Server) publish(w http.ResponseWriter, r *http.Request) {
	herald.NewHerald(s.broker)

	var body struct {
		Topic   string `json:"topic"`
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Print(err)
		return
	}

	err := s.broker.Publish(body.Topic, body.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
