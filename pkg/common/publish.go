package common

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func Publish(topic Topic, message string) {
	var data struct {
		Topic   string `json:"topic"`
		Message string `json:"message"`
	}
	data.Topic = string(topic)
	data.Message = message

	j, err := json.Marshal(&data)
	if err != nil {
		log.Print(err)
	}

	url := os.Getenv("HERALD_URI")
	http.Post(url+"/publish", "application/json", bytes.NewBuffer(j))
}
