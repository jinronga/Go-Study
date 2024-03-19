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

	//AnonymousTest()
	expressFucTest1()
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

//==============匿名函数===================

type User struct {
	id   int
	name string
}

type Manage struct {
	User
	title string
}

func (self *User) ToString() string {
	return fmt.Sprintf("id:%d,name:%s", self.id, self.name)
}

func (manage *Manage) ToString() string {
	return fmt.Sprintf("id:%d,name:%s,title:%s", manage.id, manage.name, manage.title)
}

func AnonymousTest() {
	var user User = User{1, "tom"}
	var manage Manage = Manage{user, "myTitle"}
	fmt.Println(user.ToString())
	fmt.Println(manage.ToString())
}

type T struct {
	i int
}

type S struct {
	T
}

func (t T) testT() {
	fmt.Println("类型 T 方法集包含全部 receiver T 方法。")
}

func (t *T) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
}

func testFucSet() {
	t1 := T{1}
	t2 := &t1
	fmt.Printf("t1 is : %v\n", t2)

	t2.testP()
	t2.testT()

}

func testFuncSet2() {
	t1 := T{1}
	t2 := &t1
	t1.testT()
	t1.testP()
	fmt.Printf("t1 is : %v\n", t1)
	t2.testT()
	t2.testP()
	fmt.Printf("t2 is : %v\n", t2)
}

func testFuncSet3() {
	s1 := S{T{1}}
	s2 := &s1
	fmt.Printf("s1 is : %v\n", s1)
	s1.testT()
	s1.testP()
	fmt.Printf("s2 is : %v\n", s2)
	s2.testT()
	s2.testP()
}

func (self *User) userTest() {
	fmt.Printf("%p, %v\n", self, self)
}

func expressFucTest() {

	u := User{1, "jinronga"}
	u.userTest()

	myVal := u.userTest
	myVal() // 隐式传递 receiver

	myVal2 := (*User).userTest
	myVal2(&u) // 显式传递 receiver

}

func expressFucTest1() {
	u := User{1, "Tom"}
	mValue := u.userTest // 立即复制 receiver，因为不是指针类型，不受后续修改影响。

	u.id, u.name = 2, "Jack"
	u.userTest()

	mValue()
}

func (self *User) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", self, self)
}

func (self User) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &self, self)
}

func TestPointer01(t *testing.T) {
	u := User{1, "Tom"}
	fmt.Printf("User: %p, %v\n", &u, u)

	mv := User.TestValue
	mv(u)

	mp := (*User).TestPointer
	mp(&u)

	mp2 := (*User).TestValue // *User 方法集包含 TestValue。签名变为 func TestValue(self *User)。实际依然是 receiver value copy。
	mp2(&u)
}

// ===================自定义异常=========================

func TestError01(t *testing.T) {

	//getCircleArea(5)

	FuncErrorTest02()
	fmt.Println("TestError01")
}

func FuncErrorTest() {
	a := [5]int{0, 1, 2, 3, 4}
	a[1] = 123
	fmt.Println(a)
	//a[10] = 11
	index := 10
	a[index] = 10
	fmt.Println(a)
}

func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
		// 自己抛
		panic("半径不能为负")
	}
	return 3.14 * radius * radius
}

func FuncErrorTest02() {
	// 延时执行匿名函数
	// 延时到何时？（1）程序正常结束   （2）发生异常时
	defer func() {
		// recover() 复活 恢复
		// 会返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getCircleArea(-5)
	fmt.Println("这里有没有执行")
}
