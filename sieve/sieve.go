package main

//package sieve
import (
	"fmt"
	"time"
)

func main() {
	const N = 25
	start := time.Now()
	var primeMap [N]bool
	primes := make([]int, 0)

	primeMap[0] = true
	primeMap[1] = true

	for i := 2; i < N; i++ {
		if !primeMap[i] {
			primes = append(primes, i)
			//println(i)
			for j := i * i; j < N; j += i {
				primeMap[j] = true
			}
		}
	}

	elapsed := time.Since(start)
	println(len(primes))
	fmt.Println("Primes took %s", elapsed)
}
