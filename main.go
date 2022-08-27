package main

import (
	"fmt"
	echo "github.com/milennik/restbuff/proto"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	req := &echo.EchoRequest{Name: "Nikola"}
	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatalf("Error while marshalling the object : %v", err)
	}

	res := &echo.EchoRequest{}
	err = proto.Unmarshal(data, res)
	if err != nil {
		log.Fatalf("Error while un-marshalling the object : %v", err)
	}
	fmt.Println("Value from un-marshalled data is", res.GetName())
}
