package goroutine

import (
	"fmt"
	"time"
)

func MainGoroutineChannel() {
	ch := make(chan int)

	go func() {
		ch <- 353
	}()

	val := <-ch
	fmt.Printf("got %d\n", val)

	fmt.Println("==============")
	// Send Multiple
	const count = 3
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	fmt.Println("==============")
	// Close to signal we are done
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}

	// buggered channel
	ch2 := make(chan int, 1)
	ch2 <- 19
	val2 := <-ch2
	fmt.Println(val2)
}
