package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	echo "github.com/milennik/restbuff/proto"
)

func Echo(resp http.ResponseWriter, req *http.Request) {
	contentLength := req.ContentLength
	fmt.Printf("Content Length Received : %v\n", contentLength)
	request := &echo.EchoRequest{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
		return
	}
	err = proto.Unmarshal(data, request)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
		return
	}
	name := request.GetName()
	result := &echo.EchoResponse{Message: "Hello " + name}
	response, err := proto.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	_, err = resp.Write(response)
	if err != nil {
		return
	}

}

func main() {
	fmt.Println("Starting the API server...")
	r := mux.NewRouter()
	r.HandleFunc("/echo", Echo).Methods("POST")

	server := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
