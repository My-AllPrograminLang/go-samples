// Launch some number of concurrent goroutines that sleep for a constant time,
// and measure whether the total runtime depende on the number of goroutines.
package main

import (
	"log"
	"time"
)

func launchSleepGoroutines(n int) []int {
	c := make(chan int)
	defer close(c)

	for i := 0; i < n; i++ {
		go func(i int) {
			time.Sleep(10 * time.Millisecond)
			c <- i
		}(i)
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums = append(nums, <-c)
	}

	return nums
}

func main() {
	for i := 1; i < 20; i++ {
		t1 := time.Now()
		launchSleepGoroutines(i)
		log.Printf("%2s elapsed: %s", i, time.Since(t1))
	}
}
