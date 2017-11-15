package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func sendNotifications(notifier *Notifier) {
	for {
		//every second...
		time.Sleep(time.Second)
		//send a notification message to the all WebSocket clients
		log.Printf("notifying clients with test message")
		msg := fmt.Sprintf("test message pushed to client at %v", time.Now())
		notifier.Notify([]byte(msg))
	}
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	//create a notifier and start a goroutine
	//that repeatedly sends notifications to all
	//WebSocket clients
	notifier := NewNotifier()
	go sendNotifications(notifier)

	mux := http.NewServeMux()
	mux.Handle("/websockets", NewWebSocketsHandler(notifier))

	log.Printf("server is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
