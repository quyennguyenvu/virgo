package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"virgo/service"
)

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

func respond(w http.ResponseWriter, scReturn service.Response) {
	if scReturn.Err != nil {
		http.Error(w, http.StatusText(scReturn.Code), scReturn.Code)
		return
	}

	w.WriteHeader(scReturn.Code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(scReturn.Data)
}
