package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// go func() {
	// 	msg := "this is the message"
	// 	time.Sleep(time.Second)
	// 	messages <- msg
	// }()

	time.Sleep(2 * time.Second)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
