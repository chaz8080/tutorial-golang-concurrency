package main

import "fmt"

// main listens for data from the channel(s)
func main() {
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}

// boring is generating data and sending to the channel
func f(left, right chan int) {
	left <- 1 + <-right
}
