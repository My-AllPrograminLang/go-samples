package main

import (
	"fmt"
	pb "go-samples/protobuf/person"
	"log"

	"github.com/golang/protobuf/proto"
)

func decodePerson(data []byte) {
	person := pb.Person{}
	if err := proto.Unmarshal(data, &person); err != nil {
		log.Fatal(err)
	}
	fmt.Println(person.String())
}

func main() {
	joe := &pb.Person{
		Name:   "joe",
		Emails: []string{"joe@joe.org", "joe@work.com"},
	}

	data, err := proto.Marshal(joe)
	if err != nil {
		log.Fatal(err)
	}

	decodePerson(data)
}
