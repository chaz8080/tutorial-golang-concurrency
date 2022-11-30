package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main listens for data from the channels
func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 5; i++ {
		fmt.Println(<-c) // rx an expression and print
	}
	fmt.Println("You're boring, I'm leaving")
}

// fanIn funnels 2 channels to one
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
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
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s, %d", msg, i) // send any suitable expression to the channel
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
