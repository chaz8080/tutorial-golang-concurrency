package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main listens for data from the channel(s)
func main() {
	c := boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second): // recomputed on each loop
			fmt.Println("Too slow!!!")
			return
		}
	}
}

// boring is generating data and sending to the channel
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s, %d", msg, i) // send any suitable expression to the channel
			time.Sleep(time.Duration(rand.Intn(1e3+500)) * time.Millisecond)
		}
	}()
	return c
}
