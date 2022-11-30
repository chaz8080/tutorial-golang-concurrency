package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main listens for data from the channel(s)
func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're boring, I'm leaving")
}

// fanIn funnels 2 channels to one
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	// this has the same behavior as sequencing, but only uses one go routine here instead of 2
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
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
