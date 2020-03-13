package tool

import "sync"

type wpool struct {
	Channel chan int
	Wg      *sync.WaitGroup
	Work    []func()
}

func NewPool(num int) *wpool {
	wp := new(wpool)
	wp.Wg = new(sync.WaitGroup)
	wp.Work = make([]func(), 0, 10)
	wp.Channel = make(chan int, num)
	return wp
}

func (wp *wpool) Add(callFunc func()) {
	wp.Work = append(wp.Work, callFunc)
}

func (wp *wpool) Run() {
	for k, work := range wp.Work {
		wp.Wg.Add(1)
		wp.Channel <- k
		go func(callBack func()) {
			callBack()
			defer func() {
				<-wp.Channel
				wp.Wg.Done()
			}()
		}(work)
	}
	wp.Wg.Wait()
}