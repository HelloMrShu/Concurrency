package main

import (
	"context"
	"fmt"
	"sync"
)

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

func Init(tc int) (ws *Workshop) {
	ctx, cancel := context.WithCancel(context.Background())

	ws = &Workshop{
		pipe:   make(chan Job),
		ctx:    ctx,
		cancel: cancel,
	}

	ws.hire(tc)

	return
}

func (ws *Workshop) Do(job Job) {
	ws.pipe <- job
}

func (ws *Workshop) Close() {
	defer ws.wg.Wait()
	ws.cancel()
}

func (ws *Workshop)hire(tc int)  {
	for i := 0; i < tc; i++ {
		ws.wg.Add(1)
		go func(id int) {
			defer ws.wg.Done()
			NewWorker(id, ws.pipe).Start(ws.ctx)
		}(i)
	}
}

func NewWorker(id int, pipe <-chan Job) (w *Worker) {
	w = &Worker{
		id:   id,
		pipe: pipe,
	}

	return
}

func (w *Worker) Start(ctx context.Context) {
	defer fmt.Printf("worker %d stopped\n", w.id)

	for {
		select {
		case job := <-w.pipe:
			job.Process()
		case <-ctx.Done():
			return
		}
	}
}

type SimpleJob struct {
	id int
}
func (j *SimpleJob)Process()  {
	fmt.Printf("job %d processed\n", j.id)
}

func main() {
	ws := Init(5)
	for i := 0; i < 50; i++ {
		job := &SimpleJob{id: i}
		ws.Do(job)
	}
	ws.Close()

	fmt.Println("main goroutine over")
}