package main

import "time"

func main() {

	//--	Sending and Receiving from a Channel

	// To send a value into a channel, you use the <- operator, with the channel on the left and the value to send on the right. Conversely, to receive a value from a channel, you use the <- operator with the channel on the right.

	unbufferedChannel := make(chan int)

	// Send a value into the channel
	go func() {
		unbufferedChannel <- 42
	}()

	// Receive the value from the channel
	value := <-unbufferedChannel
	println("From unbuffered channel:", value)

	//--	Buffered Channels

	// By default, channels are unbuffered, meaning that they will only hold one value and sending operations block until another goroutine receives from the channel. However, you can create a buffered channel that can hold a specified number of values before blocking.

	bufferedChannel := make(chan int, 2)

	// Now we can send two values without blocking
	bufferedChannel <- 42
	bufferedChannel <- 27

	// And receive them
	println("From buffered channel:", <-bufferedChannel)
	println("From buffered channel:", <-bufferedChannel)

	//--	Select and Timeouts
	// The select statement lets you wait on multiple channel operations, acting like a switch but for channels.

	channel := make(chan bool)
	timeout := time.After(1 * time.Second)

	go func() {
		// Simulate a task
		time.Sleep(2 * time.Second)
		channel <- true
	}()

	select {
	case <-channel:
		println("Operation successful")
	case <-timeout:
		println("Operation timeout")
	}

	//--	Closing Channels
	// You can close a channel when no more values will be sent on it.

	channel2 := make(chan int, 2)
	channel2 <- 42
	close(channel2)

	if value2, ok := <-channel2; ok {
		println("Received:", value2)
	} else {
		println("Channel closed")
	}

}
