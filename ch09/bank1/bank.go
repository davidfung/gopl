// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

import (
	"fmt"
)

type withdraw struct {
	amount int
	ch     chan bool
}

var deposits = make(chan int)        // send amount to deposit
var withdrawal = make(chan withdraw) // send amount to withdraw
var balances = make(chan int)        // receive balance

func Deposit(amount int) { deposits <- amount }
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawal <- withdraw{amount, ch}
	return <-ch
}
func Balance() int { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			fmt.Println("Depositing $", amount)
			balance += amount
		case withdraw := <-withdrawal:
			fmt.Println("Withdrawing $", withdraw.amount)
			if balance-withdraw.amount >= 0 {
				balance -= withdraw.amount
				withdraw.ch <- true
			} else {
				withdraw.ch <- false
			}
		case balances <- balance:
			fmt.Println("Balance is $", balance)
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
