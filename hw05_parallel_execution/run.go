package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var ErrNoWorker = errors.New("need at least one worker")
var ErrLimitErrors = errors.New("m should be more zero")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrLimitErrors
	}

	if n <= 0 {
		return ErrNoWorker
	}

	taskChannel := make(chan Task)
	var errCnt int32

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range taskChannel {
				if err := task(); err != nil {
					atomic.AddInt32(&errCnt, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errCnt) >= int32(m) {
			break
		}
		taskChannel <- task
	}
	close(taskChannel)

	wg.Wait()

	if errCnt >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
