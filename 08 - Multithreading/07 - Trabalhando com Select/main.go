package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	var i int64 = 0
	c1 := make(chan Message)
	c2 := make(chan Message)

	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			// time.Sleep(time.Second * 1)
			c1 <- Message{i, "Hello from RabbitMQ"}
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			// time.Sleep(time.Second * 1)
			c2 <- Message{i, "Hello from Kafka"}
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from c1: ID=%d, Msg=%s\n", msg.id, msg.Msg)
		case msg := <-c2:
			fmt.Printf("Received from c2: ID=%d, Msg=%s\n", msg.id, msg.Msg)
		case <-time.After(time.Second * 3):
			println("Timeout: No messages received within 3 seconds")
			// default:
			// println("No messages received yet")
		}
	}
}
