package main

import "fmt"

func main() {
	// 1. 基本的类型转换
	a := int(3.0)
	fmt.Println(a)
	// 在go中不支持变量间的隐式类型转换
	// 1. 变量间类型转换不支持
	var b int = 5.0 // 常量，常量到变量是会进行隐式转换的
	fmt.Println(b)
	c := 5.0
	fmt.Println("%T\n", c)

	var d int = int(c)
	fmt.Println(d)

}
