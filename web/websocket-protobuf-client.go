package main

import (
	"fmt"
	pb "go-samples/protobuf/person"
	"log"

	"github.com/golang/protobuf/proto"

	"golang.org/x/net/websocket"
)

var origin = "http://localhost/"
var url = "ws://localhost:1234/ws"

func main() {
	conn, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal("Cannot Dial:", err)
	}

	person := pb.Person{
		Name:   "Fooder",
		Emails: []string{"fooder@fooder.org", "supercool@dudez.com"},
	}
	sendData, err := proto.Marshal(&person)
	if err != nil {
		log.Fatal("Cannot Marshal:", err)
	}

	if err = websocket.Message.Send(conn, sendData); err != nil {
		log.Fatal("Couldn't send msg:", err)
	}

	var data []byte
	if err = websocket.Message.Receive(conn, &data); err != nil {
		log.Fatal("Couldn't receive:", err)
	}

	rcvPerson := pb.Person{}
	if err = proto.Unmarshal(data, &rcvPerson); err != nil {
		log.Fatal("Unmarshal:", err)
	}
	fmt.Println("Got back:", person.String())
}
