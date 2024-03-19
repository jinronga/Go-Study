package deferFunc

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

/**
  defer特性：
    1. 关键字 defer 用于注册延迟调用。
    2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
    3. 多个defer语句，按先进后出的方式执行。
    4. defer语句中的变量，在defer声明时就决定了。
*/

/**
defer用途：
    1. 关闭文件句柄
    2. 锁资源释放
    3. 数据库连接释放
*/

func whateverTest() {
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i)
	}
}

func TestDefer(t *testing.T) {
	//fmt.Println("start")
	whateverTest02()
}

func whateverTest02() {
	var whatever [5]struct{}

	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

type Test struct {
	name string
}

func (t *Test) Close01() {
	fmt.Println(t.name, "closed")
}

func Close02(t Test) {
	t.Close01()
}

func TestClose01(t *testing.T) {

	//方式一：
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer Close02(t)
	}

	//方式二：

	ts2 := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts2 {
		t2 := t
		defer t2.Close01()
	}
}

func TestFiloMain(t *testing.T) {
	fooDefer01()
}

func FILODemo(x int) {
	defer fmt.Println("a")

	defer fmt.Println("b")

	defer func() {
		fmt.Println(100 / x)
	}()

	defer fmt.Println("c")

}

func FILODemo02() {
	x, y := 10, 20

	defer func(x int) {
		fmt.Println("defer...", x, y)
	}(x) //被复制

	x += 10
	y += 20

	fmt.Printf("x:%d,y:%d\n", x, y)

}

//滥用 defer 可能会导致性能问题，尤其是在一个 "大循环" 里。

var lock = sync.Mutex{}

func lockTool() {
	lock.Lock()
	lock.Unlock()
}

func lockToolDefer() {
	lock.Lock()
	defer lock.Unlock()
}

func LockDemo() {

	func() {
		t1 := time.Now()
		for i := 0; i < 10000; i++ {
			lockTool()
		}

		elapsed := time.Since(t1)

		fmt.Println("time:", elapsed)
	}()

	func() {
		t1 := time.Now()
		for i := 0; i < 10000; i++ {
			lockToolDefer()
		}

		elapsed := time.Since(t1)

		fmt.Println("timeDefer:", elapsed)
	}()

}

func whateverTest01() {
	var whatever [5]struct{}

	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

//===========defer陷阱=============

//defer 与 closure

func fooDefer(a, b int) (i int, err error) {
	defer fmt.Println("first defer", err)
	defer func(err2 error) { fmt.Println("second defer ", err2) }(err)

	defer func() { fmt.Println("third  defer", err) }()

	if b == 0 {
		err = errors.New("division by zero")
		return
	}

	i = a / b
	return

	/**
	third defer err divided by zero!
	   second defer err <nil>
	   first defer err <nil>
	解释：如果 defer 后面跟的不是一个 closure 最后执行的时候我们得到的并不是最新的值。
	*/
}

func fooDefer01() (i int) {
	i = 0

	defer func() { fmt.Println("i=", i) }()

	return 2
}

//defer nil 函数

func testDeferNil() {
	var run func() = nil

	defer run()
	fmt.Println("run mothed...")
}

func TestNil01(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	testDeferNil()

	/**
	解释：名为 testDeferNil 的函数一直运行至结束，然后 defer 函数会被执行且会因为值为 nil 而产生 panic 异常。然而值得注意的是，run() 的声明是没有问题，因为在test函数运行完成后它才会被调用。


	*/
}

//在错误的位置使用 defer

func defHttp() error {
	res, err := http.Get("http://www.google.com")
	defer res.Body.Close()
	if err != nil {
		return err
	}

	// ..code...

	return nil
}
