package main

import (
	"fmt"
	"sync"
	"time"
)

type Work struct{}

func (w Work) Do() {
	time.Sleep(1 * time.Second)
}

func SingleWorker(ww []Work) {
	fmt.Println("SingleWorker")
	defer func(start time.Time) {
		fmt.Printf("Duration: %v\n", time.Now().Sub(start))
	}(time.Now())

	for _, w := range ww {
		w.Do()
	}
}

func ConcurrWorkers(ww []Work) {
	fmt.Println("ConcurrWorkers")
	defer func(start time.Time) {
		fmt.Printf("Duration: %v\n", time.Now().Sub(start))
	}(time.Now())

	var wg sync.WaitGroup
	for _, w := range ww {
		wg.Add(1)
		go func() {
			w.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	work := make([]Work, 5, 5)
	SingleWorker(work)
	ConcurrWorkers(work)
}
