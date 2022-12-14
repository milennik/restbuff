package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/golang/protobuf/proto"
	echo "github.com/milennik/restbuff/proto"
)

func makeRequest(request *echo.EchoRequest) *echo.EchoResponse {

	req, err := proto.Marshal(request)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	resp, err := http.Post("http://0.0.0.0:8080/echo", "application/x-binary", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &echo.EchoResponse{}
	err = proto.Unmarshal(respBytes, respObj)
	if err != nil {
		return nil
	}
	return respObj

}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		return
	}

	f, err := os.Open(cwd + "/client/client.go")
	if err != nil {
		return
	}

	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}

	encoded := base64.StdEncoding.EncodeToString(content)

	request := &echo.EchoRequest{
		Name: "Nikola",
		Data: encoded,
	}
	resp := makeRequest(request)
	fmt.Printf("Response from API is : %v\n", resp.GetMessage())
}
