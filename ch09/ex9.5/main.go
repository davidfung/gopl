package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("[ping-pong]")
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan int)

	start := time.Now()
	period := 10
	fn := func(in chan int, out chan int, start time.Time) {
		for {
			n := <-in
			if time.Since(start).Seconds() > float64(period) {
				fmt.Printf("done %d communications per second\n", n/period)
				close(done)
			}
			out <- n + 1
		}
	}

	go fn(ch1, ch2, start)
	go fn(ch2, ch1, start)
	ch1 <- 1
	<-done
}
