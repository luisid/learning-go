package main

import (
	"fmt"
	"math"
	"time"
)

var n = 1000000000
var maxConcurrency = 16

func main() {
	start := time.Now()
	primesMap := make([]bool, n, n)
	primes := sieve(int(math.Sqrt(float64(n))))

	println(len(primes), int(math.Sqrt(float64(n))))

	channel := make(chan []bool)

	for i := 0; i < maxConcurrency; i++ {
		bucket := len(primes) / maxConcurrency
		from := i * bucket
		end := (i + 1) * bucket
		if i+1 == maxConcurrency && len(primes)%maxConcurrency != 0 {
			end += len(primes) % maxConcurrency
		}
		go sievePrime(primesMap, primes[from:end], channel)
	}

	for i := 0; i < maxConcurrency; i++ {
		<-channel
	}

	for i := 2; i < n; i++ {
		if !primesMap[i] {
			primes = append(primes, i)
		}
	}

	// for i := 0; i < len(primes); i++ {
	// 	println(primes[i])
	// }
	elapsed := time.Since(start)
	println(len(primes))
	fmt.Println("Primes took %s", elapsed)
}

func sieve(n int) []int {
	primeMap := make([]bool, n, n)
	primes := make([]int, 0, 1)

	primeMap[0] = true
	primeMap[1] = true

	for i := 2; i < n; i++ {
		if !primeMap[i] {
			primes = append(primes, i)
			for j := i; j < n; j += i {
				primeMap[j] = true
			}
		}
	}

	return primes
}

func sievePrime(primeMap []bool, primes []int, channel chan []bool) {
	for i := 0; i < len(primes); i++ {
		for j := primes[i]; j < len(primeMap); j += primes[i] {
			primeMap[j] = true
		}
	}

	channel <- primeMap
}
