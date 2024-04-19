package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			// 处理数据...
			result := i
			ch <- result
			fmt.Println("Goroutine", i, "finished")
		}(i)
	}
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Result: %d\n", <-ch)
	}
	for {

	}
}
