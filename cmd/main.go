package main

import (
	"fmt"

	"github.com/mouckatron/js8call-go"
)

func main() {
	client := js8call.MakeTCPClient("localhost", "2442")
  client.Debug = true
	go client.Start()

	for {
		select {
		case message := <-client.MessageChannel:
			go handleMessage(message)
		case error := <-client.ErrorChannel:
			fmt.Println(error)
		}
	}
}

func handleMessage(message interface{}) {
	fmt.Printf("%+v\n", message)
}
