package main

import "fmt"

var a = 1

func main() {
	// 变量定义 方法一
	var a int
	a = 1
	fmt.Println(a)

	// 变量定义 方法二
	var b = 10
	fmt.Println(b)

	// 变量定义 方法三
	c := 100
	fmt.Println(c)

	// 多变量定义
	var d, e, f int
	d, e, f = 10, 20, 30
	fmt.Println(d, e, f)

	// 变量组
	var (
		num  int
		name string
	)
	fmt.Println(num, name)

	// 匿名变量（变量一旦被定义，不使用会报错的）

	// 常量（常量的数据类型可以是布尔、数字和字符串，编译时不曾使用的常量不会报错）
	const PI = 3.1415926

	// 常量组（如不指定类型，该类型和值和上一行一致）
	const (
		Female  = 0
		Male    = 1
		Unknown = 2
	)

	// 常量的枚举iota，常量的计数器 枚举, 常量等于上一个常量的表达式
	const (
		Book = iota
		Cloth
		Phone
		Food
		// iota 只能在常数组中使用，各个const定义块中互不干扰，没有表达式的常量复用上一行表达式，从第一行开始，iota从0逐渐加1
	)

	// 变量的作用域

}
