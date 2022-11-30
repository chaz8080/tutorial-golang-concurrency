package main

import (
	"fmt"
	"math/rand"
)

// main listens for data from the channel(s)
func main() {
	quit := make(chan string)
	c := boring("Joe", quit)
	for i := rand.Intn(30); i >= 0; i-- {
		fmt.Println(<-c, i)
	}
	quit <- "Bye!"
	fmt.Printf("Generator said: %q\n", <-quit)
}

// boring is generating data and sending to the channel
func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s, %d", msg, i):
				// send any suitable expression to the channel
			case <-quit:
				fmt.Println("Quitting")
				quit <- "See ya!"
				return
			}
		}
	}()
	return c
}
