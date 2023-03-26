package main

import "fmt"

func main() {
	// 很多场景下，数字有上限，我们可以选择合适的数据类型来降低内存你的占用

	// 定义bool类型
	var gender bool = true
	fmt.Println(gender)

	// int类型
	// 是一种动态类型，取决机器本身多少位，64位机器上运行，那么int就是int64，32位机器上就是int32
	var age int = 18
	fmt.Println(age)

	// float类型
	var weight float64 = 71.2
	fmt.Println(weight)
	fmt.Printf("%T\n", weight)
	// 为什么64位的float最大值远大于int64，float底层存储和int存储不一样。

	// byte rune
	// byte uint8别称，等于uint8 rune int32别称，等于int32
	var a byte = 97
	fmt.Printf("%c\n", a)
	var A rune = 65
	fmt.Printf("%c\n", A)

	// 字符
	b := 'b'
	fmt.Println("%T\n", b+1)

}
