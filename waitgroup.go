package main

import (
	"fmt"
	"sync"
	"time"
)

//waitGroup 适用于一组goroutine协同完成一个任务，当全部的goroutine都完成，任务才完成

/*
 * 例如一个服务需要调用底层多个服务的数据，那么起多个goroutine分别来调用每个服务来获取数据
 * 当全部goroutine都完成时，上层服务再统一处理各个服务的数据
*/
var wg sync.WaitGroup

func rpc1() {
	fmt.Println("rpc 1")
	wg.Done()
}

func rpc2() {
	fmt.Println("rpc 2")
	wg.Done()
}

func main() {
	wg.Add(2)

	go rpc1()
	go rpc2()

	wg.Wait()

	time.Sleep(1 * time.Second)
	fmt.Println("task is done")
}