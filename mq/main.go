package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// <-chan means this is a channel you can only read from
// cannot write to the channel
// makes sense b/c this is only listening
// chan without the error we have rw privilege
func listen(msgs <-chan amqp.Delivery) {
	log.Println("listening for new messages...")
	// can range over channels
	// once there is a message to read, this will do the things
	// in the for each
	// once it's done with the current msg body, it will wait
	// until there's another msg to act on
	for msg := range msgs {
		log.Println(string(msg.Body))
	}
}

func main() {
	mqAddr := os.Getenv("MQADDR")
	if len(mqAddr) == 0 {
		mqAddr = "localhost:5672"
	}
	mqURL := fmt.Sprintf("amqp://%s", mqAddr)
	conn, err := amqp.Dial(mqURL)
	if err != nil {
		log.Fatalf("error connecting to RabbitMQ: %v", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("error creating channel: %v", err)
	}
	q, err := channel.QueueDeclare("testQ", false, false, false, false, nil)

	// true = auto acknowledge messages
	// msgs <- chan Delivery when hover msgs
	// chan is a go channel
	msgs, err := channel.Consume(q.Name, "", true, false, false, false, nil)
	// sits and listens on own goroutine, doesn't block current one
	go listen(msgs)
	// if err != nil {
	// 	log.Fatalf("error ")
	// }

	// make chan and then the type of things you're putting into the
	// channel
	neverEnd := make(chan bool)

	// read a boolean out of the channel
	// if try to read and nothing in channel
	// the goroutine blocks until there is something in the channel
	// if try to write and channel is full
	// will block until someone reads smth out
	// 		in our case, we never write a bool into it, so it never ends
	<-neverEnd
}
