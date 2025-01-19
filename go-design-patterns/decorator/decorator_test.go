package decorator

import "testing"

func Test_decorator(t *testing.T) {
	rice := NewRice()
	t.Log(rice.Eat())

	//
	noodle := NewNoodle()
	noodle.Eat()
	t.Log("一碗干净的面条", noodle.Eat())

	//
	rice = NewFriedEggDecorator(rice)
	t.Log("米饭加个煎蛋", rice.Eat())

	noodle = NewHamSausageDecorator(noodle)
	t.Log("面条加份火腿", noodle.Eat())

	// 米饭再分别加个煎蛋和一份老干妈
	rice = NewFriedEggDecorator(rice)
	rice = NewLaoGanMaDecorator(rice)
	rice.Eat()
	t.Log("米饭加个煎蛋加份老干妈", rice.Eat())
}
