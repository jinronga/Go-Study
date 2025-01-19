package singleton

var s *singleton

type singleton struct {
}

// ====================饿汉式设计模式================================
func init() {
	s = newSingleton()
}
func newSingleton() *singleton {
	return &singleton{}
}

func (s *singleton) Running() {

}

type Instance interface {
	Running()
}

func GetInstance() Instance {
	return s
}
