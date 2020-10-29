package sync

import "sync"

type WaitGroupWrap struct {
	sync.WaitGroup
}

func (wg *WaitGroupWrap) Do(f func()) {
	wg.Add(1)
	go func() {
		f()
		wg.Done()
	}()
}
