# Concurrency
Golang concurrency control



To be honest, concurrency control feature in golang is very interesting and meaningful, and it helps us to do more efficient task, for example, download, request and so on. 

Here we will do some jobs to discuss the feature in some scenario.

### 1. sync.WaitGroup

---
Some tasks need a group of goroutine to complete and a goroutine just does a part of the task. We can use the WaitGroup to do it.
The core methods are as follows:
~~~
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