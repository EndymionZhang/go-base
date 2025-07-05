package main

import (
	"fmt"
	"time"
)

func main() {
	printNum()
	tasks := []func(){
		func() {
			time.Sleep(time.Second * 2)
			fmt.Println("任务1完成")
		},
		func() {
			time.Sleep(time.Second * 3)
			fmt.Println("任务2完成")
		},
		func() {
			time.Sleep(time.Second * 1)
			fmt.Println("任务3完成")
		},
	}
	taskHandler(tasks)
	time.Sleep(5 * time.Second)
}

func printNum() {
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()
}

func taskHandler(tasks []func()) {
	for _, task := range tasks {
		go func() {
			startTime := time.Now()
			task()
			fmt.Println("任务执行耗时为: ", time.Now().Sub(startTime))
		}()
	}
}
