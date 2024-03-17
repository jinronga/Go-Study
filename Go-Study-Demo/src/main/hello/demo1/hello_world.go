package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//fmt.Println("Hello World！")
	//fmt.Println("Println函数结尾自动换行")

	//vartest today = "星期六"
	//vartest number = 6

	//vartest str = "字符串"
	//vartest num = 1233
	//
	//str := "字符串"
	//num := 1233
	//
	//fmt.Printf("字符串：%s num：%d", str, num)
	switchDemo()
}

func conditionsDemo() {
	a := 100
	if a <= 100 {
		fmt.Println("a的值符合条件1：在0~100内,为：", a)

		if a%5 == 0 {
			fmt.Println("a的值符合条件1：在0~100内,为：", a)
		} else {
			fmt.Println("a的值不符合条件2:不可以被5整除，为：", a)
		}
	} else {
		fmt.Println("a的值不符合条件1：非在0~100内,为:", a)
	}
}

func switchDemo() {
	score := 95
	var grade string
	switch {
	case score < 60:
		grade = "C"
		fallthrough
	case false:
		fmt.Println("使用fallthrough语句跳入到了下一步打印此句,因为您的成绩为:", grade)
	case score >= 60 && score < 80:
		grade = "B"
		fallthrough
	case false:
		fmt.Println("使用fallthrough语句跳入到了下一步打印此句,因为您的成绩为:", grade)
	case score >= 80:
		grade = "A"
		fallthrough
	case false:
		fmt.Println("使用fallthrough语句跳入到了下一步打印此句,因为您的成绩为:", grade)
	}
	fmt.Println(grade)
}

func mothedTest() {

	//init 0
	var num1 int
	fmt.Println("num1 init value  is ：", num1)

	var str1 string
	fmt.Println("str1 init value  is ：", str1)

	var bool1 bool
	fmt.Println("bool1 init value  is ：", bool1)

	var a *int

	var b []int
	var c map[string]int
	var d chan int
	var e func(string2 string) int
	var f error
	fmt.Println("a,b,c,d,e,f的初始值分别为：", a, b, c, d, e, f)
}

func determineTheTypeOfTheVar() {

	str2 := "str"
	num := 123
	flag2 := true

	fmt.Println("str2:", reflect.TypeOf(str2))
	fmt.Println("num2", reflect.TypeOf(num))
	fmt.Println("flag2", reflect.TypeOf(flag2))

}

func multivariable() {

	str1, str2, str3 := "hello", "world", "aaa"

	num1, num2, num3 := 1, 2, 3
	fmt.Println(str1, str2, str3)
	fmt.Println(num1, num2, num3)

}

func referenceType() {

	var val1 = "obj"
	var val2 = 10

	fmt.Println("val1:", val1)
	fmt.Println("val2:", val2)
	fmt.Println("val1:", &val1)
	fmt.Println("val2:", &val2)
}

func constDemo() {

	//const a string = "abc" //显示类型定义
	//const b = "abc"        //隐式类型定义
	//
	//const c, d = "o", "k"
	//println(a, b, c, d)

	const (
		Unknown = 0
		Female  = 1
		Male    = 3
	)

	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)

	println(Unknown, Female, Male, a, b, c)

}
func constDemo2() {

	const (
		a = iota
		b
		c
		d = "dd" //独立值，iota += 1
		e        //值：dd
		f = 100
		g        //值：100
		h = iota //值：7 恢复计算
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

func symbolDemo() {
	a := 1
	b := 2
	fmt.Println("a+b: ", a+b)
	fmt.Println("a-b", a-b)
	fmt.Println("a*b:", a*b)
	fmt.Println("a/b：", a/b)
	fmt.Println("a%b:", a%b)

}

func symbolDemo1() {

	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c)
}

func valuationDemo2() {

	var a int = 21
	var c int

	c = a
	fmt.Println("【=】运算符 c值为=%d", c)

	c += a
	fmt.Println("【+】运算符 c值为=%d", c)
	c -= a
	fmt.Println("【-】运算符 c值为=%d", c)
	c *= a
	fmt.Println("【*】运算符 c值为=%d", c)
	c /= a
	fmt.Println("【/】运算符 c值为=%d", c)

	c = 200
	c <<= 2
	fmt.Println("【<<=】运算符 c值为=%d", c)
	c >>= 2
	fmt.Println("【>>=】运算符 c值为=%d", c)

	c &= 2
	fmt.Println("【&】运算符 c值为=%d", c)

	c ^= 2
	fmt.Println("【^】运算符 c值为=%d", c)

	c |= 2
	fmt.Println("【|】运算符 c值为=%d", c)
}
