package main

import (
	"fmt"
	"net/http"
	"virgo/config"
	"virgo/router"
	"virgo/storage"
)

func main() {
	storage.Connect()
	defer storage.Disconnect()

	r := router.NewRouter()
	serve := config.GetServe()
	fmt.Println("Listening on port " + serve.Port)
	http.ListenAndServe(":"+serve.Port, r)
}
