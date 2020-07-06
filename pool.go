package tool

import (
	"sync"
)

type workerPool struct {
	taskQueue chan func()
	capacity  int
	wg        sync.WaitGroup
}

var defaultCapacity = 5

//创建协程池
func NewWorkerPool() *workerPool {
	wp := new(workerPool)
	return wp
}

//开始并设置协程数
func (pool *workerPool) Open(capacity int) {
	if capacity == 0 {
		capacity = defaultCapacity
	}
	pool.capacity = capacity
	pool.taskQueue = make(chan func(), pool.capacity)
	pool.execTask()
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
func (pool *workerPool) Close() {
	close(pool.taskQueue)
	pool.wg.Wait()
}
