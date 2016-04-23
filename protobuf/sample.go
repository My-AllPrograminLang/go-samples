package main

import (
	"fmt"
	pb "go-samples/protobuf/person"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	joe := pb.Person{
		Name:   "joe",
		Emails: []string{"joe@joe.org", "joe@work.com"},
	}

	// Marshal and Unmarshal expect an object implementing the proto.Message
	// interface. type *Person has the interface implemented (note: a pointer to
	// Person, not just Person). Therefore here and below we always pass a pointer
	// to proto.
	data, err := proto.Marshal(&joe)
	if err != nil {
		log.Fatal(err)
	}

	decodePerson(data)
}

func decodePerson(data []byte) {
	person := pb.Person{}
	if err := proto.Unmarshal(data, &person); err != nil {
		log.Fatal(err)
	}
	fmt.Println(person.String())
}
