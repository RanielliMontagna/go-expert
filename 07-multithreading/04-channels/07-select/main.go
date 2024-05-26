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
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{id: i, Msg: "Hello from c1"}
			c1 <- msg
			// time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{id: i, Msg: "Hello from c2"}
			// time.Sleep(2 * time.Second)
			c2 <- msg
		}
	}()

	// for i := 0; i < 3; i++ {
	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("received %d: %s\n", msg1.id, msg1.Msg)

		case msg2 := <-c2:
			fmt.Printf("received %d: %s\n", msg2.id, msg2.Msg)

		case <-time.After(3 * time.Second):
			println("timeout")
			// default:
			// 	println("no one was ready")
		}
	}

}
