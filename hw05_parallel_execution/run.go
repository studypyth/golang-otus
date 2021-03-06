package hw05

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type Cnt struct {
	mu  sync.Mutex
	cnt int
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, gortnNmb, errLimit int) error {
	taskChan := make(chan Task, len(tasks))
	guard := make(chan struct{}, gortnNmb)
	errCnt := Cnt{}
	wg := sync.WaitGroup{}
	var res error
	for _, task := range tasks { // Выдаем задания в канал тасков для воркеров и закрываем канал - заданий больше нет
		if errCnt.cnt == errLimit {
			res = ErrErrorsLimitExceeded
			break
		}
		taskChan <- task
		wg.Add(1)
		guard <- struct{}{}
		go func(task Task) {
			worker(task, &wg, &errCnt)
			<-guard
		}(task)
	}
	wg.Wait()
	return res
}
func worker(task Task, wg *sync.WaitGroup, errCnt *Cnt) {
	defer wg.Done()
	res := task()
	if res != nil {
		errCnt.mu.Lock()
		errCnt.cnt++
		errCnt.mu.Unlock()
	}
}
