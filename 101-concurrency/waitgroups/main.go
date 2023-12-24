package main

import (
	"log"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	defer log.Printf("Worker #%d done", id)

	log.Printf("Worker #%d starting", id)
	time.Sleep(time.Second)
}

func main() {
	// This WaitGroup is used to wait for all the goroutines launched here to finish. Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each.
	for i := 1; i <= 500; i++ {
		id := i

		wg.Add(1)
		go worker(&wg, id)
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
	wg.Wait()

	log.Printf("All workers done")
}
