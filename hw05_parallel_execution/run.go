package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	var errCount uint64
	var mutex sync.Mutex

	if m < 0 {
		m = 0
	}

	ch := make(chan Task, len(tasks))

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			for {
				t, ok := <-ch
				mutex.Lock()
				errCountCp := errCount
				mutex.Unlock()

				if !ok || errCountCp >= uint64(m) {
					wg.Done()
					return
				}
				err := t()
				if err != nil {
					mutex.Lock()
					errCount++
					mutex.Unlock()
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
