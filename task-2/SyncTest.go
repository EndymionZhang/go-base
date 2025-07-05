package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var lock = sync.Mutex{}

func main() {
	count := 0
	for i := 0; i < 10; i++ {
		go counter(&count)
	}
	time.Sleep(time.Second * 5)
	fmt.Println(count)

	count2 := atomic.Int64{}
	count2.Store(0)
	for i := 0; i < 10; i++ {
		go counter2(&count2)
	}
	time.Sleep(time.Second * 5)
	fmt.Println(count2.Load())
}

func counter(count *int) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count++
		lock.Unlock()
		fmt.Println(*count)
	}
}

func counter2(count *atomic.Int64) {
	for i := 0; i < 1000; i++ {
		// count 加一
		count.Add(1)
		fmt.Println(count.Load())
	}

}
