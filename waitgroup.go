package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := "https://www.sohu.com"
	wg := &sync.WaitGroup{}
	for i:= 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			fmt.Println("goroutine: ", i)
			client := http.Client{Timeout: 20 * time.Second}
			resp, err := client.Get(url)
			if err != nil {
				fmt.Println("request error", err, resp)
			}
		}(i, url)
	}
	time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println("wait_group is over")
}