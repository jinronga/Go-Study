package deferFunc

import (
	"fmt"
	"io"
	"net/http"
	"os"
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

func TestDoHttp(t *testing.T) {
	openFile03()
}

func doGet() error {
	res, err := http.Get("http://www.google.com")

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return err
	}

	return nil
}

func openFile() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}

	if f != nil {
		defer func() {
			if err := f.Close(); err != nil {
				// log etc
			}
		}()
	}

	// ..code...

	return nil
}

func openFile01() (err error) {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}

	if f != nil {
		defer func() {
			if ferr := f.Close(); ferr != nil {
				// log etc
				err = ferr
			}
		}()
	}

	// ..code...

	return nil
}

func openFile02() (err error) {

	file, err := os.Open("book.text")

	if err != nil {
		return err
	}

	if file != nil {
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}()
	}

	file, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	if file != nil {
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}()
	}
	return nil
}

/**
当延迟函数执行时，只有最后一个变量会被用到，因此，f 变量 会成为最后那个资源 (another-book.txt)。而且两个 defer 都会将这个资源作为最后的资源来关闭
*/
//解决方案：

func openFile03() (err error) {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}(f)
	}

	// ..code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}(f)
	}

	return nil
}
