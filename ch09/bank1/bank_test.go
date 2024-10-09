// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"
	"time"

	bank "bank1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("Balance =", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	time.Sleep(1 * time.Second)
	// Candy
	go func() {
		if bank.Withdraw(50) {
			fmt.Println("withdraw successful")
		} else {
			fmt.Println("withdraw failed")
		}
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	<-done

	if got, want := bank.Balance(), 250; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
