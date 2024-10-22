package main

import (
	"fmt"
	"testing"
	"time"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestMe(t *testing.T) {
	fmt.Println("pipeline...")

	first := make(chan int)
	last := make(chan int)

	var in chan int
	var out chan int
	out = first

	stages := 1_600_000
	fmt.Printf("stages=%d\n", stages)

	start := time.Now()
	for i := 1; i <= stages; i++ {
		// fmt.Printf("creating goroutine #%d\n", i)
		in = out
		if i == stages {
			out = last
		} else {
			out = make(chan int)
		}
		go func(in chan int, out chan int) {
			n := <-in
			// fmt.Printf(("incrementing %d to %d...\n"), n, n+1)
			n++
			out <- n
		}(in, out)
	}
	fmt.Printf("%d stages created in %v secs\n", stages, time.Since(start))

	start = time.Now()
	first <- 0
	<-last
	fmt.Printf("%d stages passed thru in %v secs\n", stages, time.Since(start))
}
