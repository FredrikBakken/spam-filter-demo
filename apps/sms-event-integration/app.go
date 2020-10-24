package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"telenor.com/spam-filter-demo/sms-event-integration/config"
	"telenor.com/spam-filter-demo/sms-event-integration/handlers"
)

func main() {
	fmt.Println("Starting the SMS Event Integration application...")

	cfg := config.New()
	go handlers.NewMessageEvent(cfg)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(cfg.Server.Route, handlers.NewSmsEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, router))
}
