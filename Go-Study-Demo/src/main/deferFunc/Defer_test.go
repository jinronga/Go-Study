package deferFunc

import (
	"fmt"
	"testing"
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
