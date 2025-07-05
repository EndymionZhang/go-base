package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(<-ch)
		}
	}()

	ch1 := make(chan int, 100)
	go chanReceive(ch1)
	go chanSend(ch1)

	time.Sleep(time.Second * 5)
}

func chanSend(sender chan<- int) {
	for i := 1; i <= 100; i++ {
		sender <- i
	}
}

func chanReceive(receiver <-chan int) {
	for i := 1; i <= 100; i++ {
		fmt.Println("receiver", <-receiver)
	}
}
