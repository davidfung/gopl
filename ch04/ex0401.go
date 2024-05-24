// Ex 14.1
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Ex 4.1")
	// sha256 hash
	//                        1               2               3
	//        0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
	hash1 := "9834876dcfb05cb167a5c24953eba58c4ac89b1adf57f28f2f9d09af107ee8f0"
	// hash2 := "8834876dcfb05cb167a5c24853eba58c4ac88b2adf57f28f2f9d09af107ee8f1"
	hash2 := "3cf9a1a81f6bdeaf08a343c1e1c73e89cf44c06ac2427a892382cae825e7c9c1"
	cnt := countDiffBits(hash1, hash2)
	fmt.Printf("hash1 = %s\nhash2 = %s\ndiff bits = %d\n", hash1, hash2, cnt)
}

func countDiffBits(hash1, hash2 string) int {
	//todo
	h1 := []string{hash1[:16], hash1[16:32], hash1[32:48], hash1[48:]}
	h2 := []string{hash2[:16], hash2[16:32], hash2[32:48], hash2[48:]}
	count := 0
	for i := range h1 {
		fmt.Println(i)
		fmt.Println(h1[i])
		fmt.Println(h2[i])
	    u1, _ := strconv.ParseUint(h1[i], 16, 64)
	    u2, _ := strconv.ParseUint(h2[i], 16, 64)
	    u := u1 ^ u2
	    count += PopCount(u)
	}
	return count
}

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}