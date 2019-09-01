package main

import (
	"fmt"
	"net/http"
	"virgo/config"
	"virgo/helper"
	"virgo/router"
	"virgo/storage"
)

func main() {
	// connect database
	storage.Connect()
	defer storage.Disconnect()

	// register logger
	helper.Logging()
	defer helper.CloseFile()

	r := router.NewRouter()
	serve := config.GetServe()
	fmt.Println("Listening on port " + serve.Port)
	http.ListenAndServe(":"+serve.Port, r)
}
