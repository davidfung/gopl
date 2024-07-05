package main

import (
	"fmt"
	"time"
)

func main() {
	longRunningFunc()
}

func longRunningFunc() {
	defer trace("longRunningFunc()")()
	time.Sleep(time.Duration(2) * time.Second)
}

func trace(name string) func() {
	var t = time.Now()
	fmt.Printf("Entering %s at %s\n", name, t)
	return func() {
		fmt.Printf("Exiting %s taken %s \n", name, time.Now().Sub(t))
	}
}
