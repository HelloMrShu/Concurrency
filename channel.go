package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	limit := 3
	wg := sync.WaitGroup{} //to make sure all goroutines are done
	ch := make(chan int, limit) // channel with buffer and blocked when buffer is full
	defer close(ch)

	for i := 0; i < 30; i++ {
		ch <- i
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("goroutine %d: time %d\n", x, time.Now().Unix())
			test()
			time.Sleep(time.Second)
			<- ch
		}(i)
	}
	wg.Wait()
	fmt.Println("main goroutine is done")
}

func test() {
	fmt.Println("go run test func")
}
