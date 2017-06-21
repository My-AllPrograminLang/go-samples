// Performs a CPU-bound computation (factorizing large numbers) in serial and
// concurrently.
package main

import (
	"log"
	"time"
)

func factorize(n int64) []int64 {
	var factors []int64
	var p int64 = 2
	for {
		if n == 1 {
			return factors
		}
		r := n % p
		if r == 0 {
			factors = append(factors, p)
			n = n / p
		} else if p*p >= n {
			factors = append(factors, n)
			return factors
		} else if p > 2 {
			p += 2
		} else {
			p += 1
		}
	}
}

type FactorMap map[int64][]int64

func serialFactorize(nums []int64) FactorMap {
	fm := make(FactorMap)
	for _, n := range nums {
		fs := factorize(n)
		fm[n] = fs
	}
	return fm
}

func concurrentFactorize(nums []int64) FactorMap {
	c := make(chan []int64)
	defer close(c)

	for _, n := range nums {
		go func(n int64) {
			c <- factorize(n)
		}(n)
	}

	fs := make(FactorMap)
	for _, n := range nums {
		fs[n] = <-c
	}
	return fs
}

func main() {
	var n int64 = 4477457 * 982451653
	var is []int64
	for i := 0; i < 30; i++ {
		is = append(is, n)
	}

	{
		t1 := time.Now()
		fs := serialFactorize(is)
		log.Printf("%d elapsed: %s", len(fs), time.Since(t1))
	}

	{
		t1 := time.Now()
		fs := concurrentFactorize(is)
		log.Printf("%d elapsed: %s", len(fs), time.Since(t1))
	}
}
