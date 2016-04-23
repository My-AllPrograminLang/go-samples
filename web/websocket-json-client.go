package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

type Person struct {
	Name   string
	Emails []string
}

var origin = "http://localhost/"
var url = "ws://localhost:1234/ws"

func main() {
	conn, err := websocket.Dial(url, "", origin)
	checkError(err)

	person := Person{
		Name:   "Fooder",
		Emails: []string{"fooder@fooder.org", "supercool@dudez.com"},
	}

	err = websocket.JSON.Send(conn, person)
	if err != nil {
		fmt.Println("Couldn't send msg " + err.Error())
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
