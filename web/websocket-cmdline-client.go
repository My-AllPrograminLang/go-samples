package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

var origin = "http://localhost/"
var url = "ws://localhost:1234/ws"

func main() {
	// Need to set the origin to localhost, otherwise the server's default handler
	// will reject it.
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	msg := "hello there server"
	if err = websocket.Message.Send(ws, msg); err != nil {
		log.Fatal(err)
	}

	var reply string
	if err = websocket.Message.Receive(ws, &reply); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Received back:", reply)
}
