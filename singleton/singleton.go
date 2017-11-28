package singleton

import "sync"

type singleton struct {
	count int
	sync.RWMutex
}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}

func (s *singleton) AddOne() {
	s.Lock()
	defer s.Unlock()

	s.count++
}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}

var addCh chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)

func init() {
	var count int
	go func(addCh <-chan bool, geCtountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				break
			}
		}
	}(addCh, getCountCh, quitCh)
}
