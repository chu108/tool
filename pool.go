package tool

import (
	"sync"
)

type workerPool struct {
	taskQueue chan func()
	capacity  int
	wg        sync.WaitGroup
}

//创建携程池
func NewWorkerPool(capacity int) *workerPool {
	wp := new(workerPool)
	wp.capacity = capacity
	wp.taskQueue = make(chan func(), wp.capacity)

	wp.execTask()

	return wp
}

//执行任务
func (pool *workerPool) execTask() {
	for i := 0; i < pool.capacity; i++ {
		pool.wg.Add(1)
		go func() {
			defer func() {
				pool.wg.Done()
			}()
			for {
				select {
				case fn, ok := <-pool.taskQueue:
					if !ok {
						return
					}
					if fn != nil {
						fn()
					}
				}
			}
		}()
	}
}

//添加任务
func (pool *workerPool) Add(fn func()) {
	pool.taskQueue <- fn
}

//关闭任务
func (pool *workerPool) Run() {
	close(pool.taskQueue)
	pool.wg.Wait()
}
