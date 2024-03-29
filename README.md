# Concurrency
Golang concurrency control



Concurrency control feature in golang is very interesting and meaningful, and it helps us to do more efficient task, for example, download, request and so on. Here we will do some jobs to discuss the feature in some scenario.

### 1. sync.WaitGroup

---
Some tasks need a group of goroutine to complete and a goroutine just does a part of the task. We can use the WaitGroup to do that.However, wg can't control the certain number of concurrency. So you should understand the feature to avoid some mistakes.

The core methods as follows:

~~~go
wg := sync.WaitGroup
for i:=0; i < 10; i++ {
    wg.Add(1)
    go func() {
        //do something...
        wg.Done()
    }
    wg.Wait()
    //do something...
}
~~~

### 2.channel

---
The buffered channel can be used as a tool for concurrency control. And the buffer size is the concurrency limit.we can use sync.WaitGroup to make sure all goroutines are done.

```go
limit := 10 //concurrency limit is 10
wg := sync.WaitGroup
ch := make(chan int, limit)
for i:= 0; i < 10; i++ {
	wg.Add(1)
	ch <- i
	go func() {
		defer wg.Done()
		//do something
		<- ch
    }
	wg.Wait()
	fmt.Println("main goroutine is done")
}
```

### 3.channel + context
---
context controls a group of goroutine, and each has the same context, u can use the root context to handle the child context.

```go
type Workshop struct {
	pipe   chan Job
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}
type Worker struct {
	id   int
	pipe <-chan Job
}
type Job interface {
	Process()
}
```