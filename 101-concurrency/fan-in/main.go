package main

import (
	"fmt"
	"sync"
	"time"
)

// payload contains a name and a value.
type payload struct {
	name  string
	value int
}

// producer takes a name, a done channel, and a wait group
// It returns a channel of payloads
func producer(name string, done <-chan struct{}, wg *sync.WaitGroup) <-chan payload {
	ch := make(chan payload)
	var i = 1
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				close(ch)
				fmt.Println(name, "completed")
				return
			case ch <- payload{
				name:  name,
				value: i,
			}:
				fmt.Println(name, "produced", i)
				i++
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
	return ch
}

// consumer takes a slice of channels, a done channel, a wait group, and a fan-in channel
// It starts a goroutine for each channel in the slice to consume the payloads
func consumer(channels []<-chan payload, done <-chan struct{}, wg *sync.WaitGroup, fanIn chan<- payload) {
	for i, ch := range channels {
		i := i + 1
		ch := ch
		go func() {
			defer wg.Done()
			fmt.Println("Started consume of producer", i)
			for {
				select {
				case <-done:
					fmt.Println("Consume of producer", i, "completed")
					return
				case v := <-ch:
					fmt.Println("Consumer of producer", i, "got value", v.value, "from", v.name)
					fanIn <- v
				}
			}
		}()
	}
}

func main() {
	done := make(chan struct{}) // A channel used to signal when the program should stop
	wg := sync.WaitGroup{}      // A wait group used to wait for all goroutines to finish

	wg.Add(2) // Add 2 to the wait group, since we will start 2 producer goroutines

	producers := make([]<-chan payload, 0, 3)                       // Create a slice of channels to hold the producer channels
	producers = append(producers, producer("Producer1", done, &wg)) // Start Producer1 and add its channel to the slice.
	producers = append(producers, producer("Producer2", done, &wg)) // Start Producer2 and add its channel to the slice.

	fanIn := make(chan payload) // Create a channel to hold the payloads from all producers

	wg.Add(2)                             // Add 2 to the wait group, since we will start 2 consumer goroutines
	consumer(producers, done, &wg, fanIn) // Start the consumer goroutines

	go func() {
		for {
			fanInDone := false
			select {
			case <-done:
				if !fanInDone {
					continue
				}
				return
			case v, ok := <-fanIn:
				if !ok {
					fanInDone = true
				}
				fmt.Printf("fanIn got %v\n", v)
			}
		}
	}()

	time.Sleep(time.Second) // Sleep for 1 second to allow the producers and consumers to run
	close(done)             // Close the done channel to signal that the program should stop
	wg.Wait()               // Wait for all goroutines to finish
}
