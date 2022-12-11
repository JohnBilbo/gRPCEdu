package main

import (
	"fmt"
	"gRPCEdu/person"

	"github.com/golang/protobuf/proto"
)

func main() {
	newReq := &person.NewsRequest{
		Name: "Hello",
	}
	raw, errMarsh := proto.Marshal(newReq)
	if errMarsh != nil {
		fmt.Println("errMarsh", errMarsh)
	}
	fmt.Println(string(raw))
}
