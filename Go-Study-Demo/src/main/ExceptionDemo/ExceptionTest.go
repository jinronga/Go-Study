package ExceptionDemo

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {

}
func panicTest() {

	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}

	}()
	panic("panic error!")
}

func panicTest01() {

	defer func() {
		if error := recover(); error != nil {
			fmt.Println(error)
		}
	}()

	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1
}

func panicTest02() {

	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("panic test")
}

func panicTest03() {
	defer func() {
		fmt.Println(recover()) //有效
	}()
	defer recover()              //无效！
	defer fmt.Println(recover()) //无效！
	defer func() {
		func() {
			println("defer inner")
			recover() //无效！
		}()
	}()

	panic("test panic")
}

func panicTest04() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}

func TestPanic(t *testing.T) {
	//panicTest()
	//panicTest01()
	panicTest02()
}
