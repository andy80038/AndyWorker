package main

import (
	"fmt"
	"sync"
)

func main() {
	createDemo()

}
func doWork(id int, c chan int, w worker) {
	//方法二
	for ch := range c {
		fmt.Printf("Worker %d received %d \n", id, ch)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w.in, w)

	return w
}
func createDemo() {
	var workers [10]worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		wg.Add(1)
		worker.in <- i + 900

	}
	for i, worker := range workers {
		wg.Add(1)
		worker.in <- i

	}
	wg.Wait()
	//time.Sleep(time.Millisecond)
}
