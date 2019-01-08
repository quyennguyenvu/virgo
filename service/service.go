package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

// Response ..
type Response struct {
	Data []byte
	Code int
	Err  error
}

func getBody(o interface{}, readClose io.ReadCloser) {
	body, err := ioutil.ReadAll(readClose)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, o)
	if err != nil {
		log.Fatal(err)
	}
}
