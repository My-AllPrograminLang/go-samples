// TODO: real serveHtml here with JS that does the protobuf thing
package main

import (
	"fmt"
	pb "go-samples/protobuf/person"
	"log"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/golang/protobuf/proto"
)

func decodePerson(data []byte) {
	person := pb.Person{}
	if err := proto.Unmarshal(data, &person); err != nil {
		log.Fatal(err)
	}
	fmt.Println(person.String())
}

func receivePerson(ws *websocket.Conn) {
	var data []byte
	err := websocket.Message.Receive(ws, &data)
	if err != nil {
		fmt.Println("Can't receive")
	} else {
		person := pb.Person{}
		if err := proto.Unmarshal(data, &person); err != nil {
			log.Fatal(err)
		}
		log.Printf("Received person: %s", person.String())

		// Tweak the name a bit and pong it back to the client.
		person.Name = "!!" + person.Name
		sendData, err := proto.Marshal(&person)
		if err != nil {
			fmt.Println("Cannot Marshal:", err)
		}
		if err := websocket.Message.Send(ws, sendData); err != nil {
			fmt.Println("Couldn't send msg " + err.Error())
		}
	}
}

func serveHtml(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `hello`)
}

func main() {
	http.HandleFunc("/", serveHtml)
	http.Handle("/ws", websocket.Handler(receivePerson))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
