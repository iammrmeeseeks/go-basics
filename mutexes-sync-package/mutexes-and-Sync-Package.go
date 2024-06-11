package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	count++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Count:", count)
}

/*


	declares the mutex mu and the count variable.

	The increment function now accepts a sync.WaitGroup pointer wg.
		defer wg.Done() ensures that the wait group counter is decremented when the function exits.
		mu.Lock() and mu.Unlock() ensure that count++ is thread-safe.

	A sync.WaitGroup wg is created to keep track of all the goroutines.
		wg.Add(1) increments the wait group counter for each goroutine.
		wg.Wait() waits for all goroutines to finish before printing the final count.

	By using a sync.WaitGroup, we ensure that the main function waits for all the
		increment goroutines to complete before printing the final count,
		making the program more reliable and correct.

*/
