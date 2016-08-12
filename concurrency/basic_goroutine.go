package main

import (
	"fmt"
	"time"
)

/*
 * 1. Goroutines: light-weighted threads maintained by Go runtime
 */

func simple_say_routine (s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Observe that non-determinant order of printing
// Caused by arbitrary execution of 2 threads(goroutine)
func test_simple_say_routine () {
	go simple_say_routine("world")
	simple_say_routine("hello")
}

/*
 * 2. Channels: a typed conduit through which you can send and receive values
 */

// By default, sends and receives block until the other side is ready. 
// This allows goroutines to synchronize without explicit locks or 
// condition variables.

func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v

	}
	c <- sum
}

func test_sum_channel() {
	arr := []int{1,3,5,2,4,6}
	ch := make(chan int)

	go sum(arr[len(arr)/2:], ch)
	go sum(arr[:len(arr)/2], ch)

	x, y := <-ch, <-ch

	fmt.Println("one half: %d", x)
	fmt.Println("another half: %d", y)
	fmt.Println("total sum: %d", x+y) 
}

/*
 * 3. Buffered Channels: Channels can be buffered. And Channels can be closed.
 */

 // Sends to a buffered channel block only when 
 // the buffer is full. Receives block when the buffer is empty.

 // A sender can close a channel to indicate that no more values will be sent.
 // Receivers can test whether a channel has been closed.
 // Only the sender should close a channel, never the receiver.
 // Sending on a closed channel will cause a panic.
 // While channels are not supposed to be closed necessarily. Only apply close()
 // when it's needed to tell the receiver to terminate some operation

func fibonacci(ch chan int) {
	x, y := 0, 1
	for i := 0; i < cap(ch); i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func test_fibonacci_channel(n int) {
	ch := make(chan int, n)
	go fibonacci(ch)
	for i := range ch {
		fmt.Println(i)
	}
}

/*
 * 4. Select: lets a goroutine wait on multiple communication operations.
 */

// A select blocks until one of its cases can run, then it executes 
// that case. It chooses one at random if multiple are ready.

func fibonacci_select(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
			case ch <- x :
				x, y = y, x+y
			case <- quit:
				fmt.Println("Fiboniacci quit!")
				return ;		
		}
	}
}

func run_fibonacci_select() {
	ch := make(chan int)
	quit := make(chan int)

	go func(){
		for i:= 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fibonacci_select(ch, quit)
}

/*
 *Entry to the test code
 */
// func main() {
// 	//1 : goroutine
// 	//test_simple_say_routine()
// 	//2 : simple channel
// 	//test_sum_channel()
// 	//3	: channel that can be closed
// 	//test_fibonacci_channel(5)
// 	//4 : select statement usage
// 	//run_fibonacci_select()
// }
