package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	msg  string
	wait chan bool
}

// main listens for data from the channel(s)
func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		msg1 := <-c
		fmt.Println(msg1.msg)

		msg2 := <-c
		fmt.Println(msg2.msg)

		<-msg1.wait // reset channel, stop blocking
		<-msg2.wait
	}
	fmt.Println("You're boring, I'm leaving")
}

// fanIn funnels 2 channels to one
func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

// boring is generating data and sending to the channel
func boring(msg string) <-chan Message {
	c := make(chan Message)
	blockingStep := make(chan bool) // channel within channel to control exec, set false default
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s, %d", msg, i), blockingStep} // send any suitable expression to the channel
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			blockingStep <- true
		}
	}()
	return c
}
