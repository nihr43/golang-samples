package main

import (
	"context"
	"fmt"
	"time"
)

func runClock(clock chan int) {
	for i := 0; ; i++ {
		clock <- i
		time.Sleep(1 * time.Second)
	}
	close(clock)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clock := make(chan int)
	go runClock(clock)

	for {
		select {
		case i := <-clock:
			if i%2 == 0 {
				fmt.Println("tick")
			} else {
				fmt.Println("tock")
			}
		case <-ctx.Done():
			fmt.Println("context timeout met")
			return
		}
	}
}
