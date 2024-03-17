package funcDemo

import (
	"fmt"
	"testing"
)

func TestLoadFromFiles(t *testing.T) {

	//n, res := test(1, 2, "aa")
	//t.Logf("%d,%s", n, res)

	//n := testFn(func() int {
	//	return 100
	//})
	//
	//n1 := formatFuncTest(func(s string, x, y int) string {
	//	return fmt.Sprintf("%s,%d,%d", s, x, y)
	//}, "hello", 1, 2)
	//
	//println(n, n1)

	SwapTest(t)
}

func test(x, y int, s string) (int, string) {

	n := x + y
	return n, fmt.Sprintf(s, n)
}

func testFn(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string

func formatFuncTest(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func swap(x, y *int) {
	var temp int

	temp = *x /* 保存 x 的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y*/

	fmt.Sprintf("*x:%d,*y:%d,temp:%d", *x, *y, temp)
}

func SwapTest(t *testing.T) {
	var a, b int = 1, 2
	/*
	   调用 swap() 函数
	   &a 指向 a 指针，a 变量的地址
	   &b 指向 b 指针，b 变量的地址
	*/
	swap(&a, &b)
	t.Logf("a = %d,b = %d", a, b)
}
