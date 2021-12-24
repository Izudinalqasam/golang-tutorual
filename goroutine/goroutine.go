package goroutine

import (
	"fmt"
	"time"
)

// channel by default sends and receive block until the other side is ready
func MainGoroutine() {
	bufferChannel()
}

func channel() {
	ch := make(chan int)

	go hello(&ch)
	fmt.Println(time.Now(), "Waiting for message")

	func(message string) {
		fmt.Println(time.Now(), message)
	}("Waiting for message")

	v := <-ch
	fmt.Println(time.Now(), "Received message of ", v)
}

func hello(ch *chan int) {
	fmt.Println(time.Now(), "Taking a nap")

	time.Sleep(2 * time.Second)

	*ch <- 20
}

func bufferChannel() {
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(time.Now(), i, "Sending Message")
			ch <- i
			fmt.Println(time.Now(), i, "Message Sent")
		}

		fmt.Println(time.Now(), "all Completed")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Waiting for message")

	for v := range ch {
		fmt.Println(time.Now(), v, "Received Message")
	}

	fmt.Println("Exiting")
}

// fmt.Println(time.Now(), <-ch, "Received Message")
// fmt.Println(time.Now(), <-ch, "Received Message")
// fmt.Println(time.Now(), <-ch, "Received Message")
