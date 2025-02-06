package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {
	type Message struct {
		Message string `json:"message"`
	}

	message := &Message{}

	for range time.Tick(time.Second * 5) {
		client := resty.New()
		_, err := client.R().SetResult(&message).Get("http://gin-server-clusterip-service:80")
		if err != nil {
			s := fmt.Sprintf("api call failed %v", err)
			panic(s)
		}

		fmt.Println(message.Message)
	}
}
