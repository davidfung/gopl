// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(ch chan string, shout string, delay time.Duration) {
	ch <- fmt.Sprintf("\t%s", strings.ToUpper(shout))
	time.Sleep(delay)
	ch <- fmt.Sprintf("\t%s", shout)
	time.Sleep(delay)
	ch <- fmt.Sprintf("\t%s", strings.ToLower(shout))
}

func scan(c net.Conn, ch chan string) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(ch, input.Text(), 1*time.Second)
	}
}

func handleConn(c net.Conn) {
	ch := make(chan string)
	go scan(c, ch)
	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case s := <-ch:
			fmt.Fprintln(c, s)
			timer = time.NewTimer(10 * time.Second)
		case <-timer.C:
			fmt.Fprintln(c, "timeout!")
			goto Exit
		}
	}
Exit:
	fmt.Println("closing connection...")
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		fmt.Println("connection accepted")
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
