package js8call

import (
	"bufio"
	"encoding/json"

	"errors"
	"fmt"
	"net"
)

type TCPClient struct {
	Address        string
	Port           string
	MessageChannel chan interface{}
	ErrorChannel   chan error
	client         net.Conn
}

/*
MakeTCPClient creates a TCPClient with default channels
*/
func MakeTCPClient(address string, port string) *TCPClient {

	messageChannel := make(chan interface{}, 5)
	errorChannel := make(chan error, 5)

	tcpclient := &TCPClient{
		address,
		port,
		messageChannel,
		errorChannel,
		nil}

	return tcpclient
}

func (client *TCPClient) Start() { //messageChannel chan interface{}, errorChannel chan error) {

	var err error

	(*client).client, err = net.Dial("tcp", client.Address+":"+client.Port)
	if err != nil {
		fmt.Println(err)
	}

	for {
		message, _ := bufio.NewReader((*client).client).ReadString('\n')

		go client.parse(message)
	}
}

func (client *TCPClient) parse(message string) {
	var genericMessage Message
	json.Unmarshal([]byte(message), &genericMessage)

	switch genericMessage.Type {

	case "RX.DIRECTED":
		var jsonMessage RXDirected
		json.Unmarshal([]byte(message), &jsonMessage)
		// log.Println(jsonMessage)
		client.MessageChannel <- jsonMessage

	case "RX.SPOT":
		var jsonMessage RXSpot
		json.Unmarshal([]byte(message), &jsonMessage)
		// log.Println(jsonMessage)
		client.MessageChannel <- jsonMessage

	case "RX.ACTIVITY":
		var jsonMessage RXActivity
		json.Unmarshal([]byte(message), &jsonMessage)
		// log.Println(jsonMessage)
		client.MessageChannel <- jsonMessage

	case "TX.FRAME":
		var jsonMessage TXFrame
		json.Unmarshal([]byte(message), &jsonMessage)
		// log.Println(jsonMessage)
		client.MessageChannel <- jsonMessage

	case "RIG.PTT":
		var jsonMessage RigPTT
		json.Unmarshal([]byte(message), &jsonMessage)
		// log.Println(jsonMessage)
		client.MessageChannel <- jsonMessage

	default:
		// log.Println("Could not parse message", message)
		client.ErrorChannel <- errors.New("Could not parse message " + message)
	}
}
