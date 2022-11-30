package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main listens for data from the channel
func main() {
	// note: boring is acting as a service here, we could create many more of the same service
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // rx an expression
	}
	fmt.Println("You're boring, I'm leaving")
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
