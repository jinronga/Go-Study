package singleton

import "sync"

var (
	instanceLazy *singletonLazy
	mux          sync.Mutex
	once         sync.Once
)

type singletonLazy struct {
}

func newSingletonLazy() *singletonLazy {
	return &singletonLazy{}
}

type SingletonLazy interface {
	Running()
}

func (s *singletonLazy) Running() {

}

func GetSingletonLazy() SingletonLazy {
	if s != nil {
		return s
	}
	mux.Lock()
	defer mux.Unlock()
	if s != nil {
		return s
	}
	s = newSingleton()
	return s
}

func GetSingletonLazy1() Instance {
	once.Do(func() {
		s = newSingleton()
	})
	return s
}
