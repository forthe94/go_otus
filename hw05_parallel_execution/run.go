package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m < 0 {
		m = 0
	}
	var wg sync.WaitGroup
	var errCount uint64

	ch := make(chan Task, len(tasks))

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			for {
				t, ok := <-ch
				if !ok || errCount >= uint64(m) {
					wg.Done()
					return
				}
				err := t()
				if err != nil {
					atomic.AddUint64(&errCount, 1)
				}
			}
		}()
	}

	for _, v := range tasks {
		ch <- v
	}
	close(ch)
	wg.Wait()
	if errCount >= uint64(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
