package main

import "fmt"

// Programmer 接口是一个协议，程序员——只要你能够写代码、解决bug，其实就是一组方法的集合
type Programmer interface {
	Coding() string
	Debug() string
}

type Pythoner struct {
	yeears int
	lib    []string
	kj     []string
}

func (p Pythoner) Coding() string {
	fmt.Println("Python开发者")
	return "Python开发者"
}

func (p Pythoner) Debug() string {
	fmt.Println("我会Python的debug")
	return "我会Python的debug"
}

// 开发中经常会遇到的问题，开发一个电商网站，支付环节，使用微信、支付宝银行卡
// 定义一个协议 1.创建订单 2.支付 3.查询支付状态 4.退款
// 支付发起了

type AliPay struct {
}

type WechatPay struct {
}

type BankPay struct {
}

var a = AliPay{}
var b = BankPay{}
var w = WechatPay{}

// 如果后期接入一个新的支付，或者取消已有支付，需要改代码

// UnionPay 多态 声明一种类型的时候声明一种兼容类型，但是实际赋值的时候是另一种类型
// 接口的强制性
var UnionPay interface {
}

// go语言中不支持继承

func main() {
	var pro Programmer = Pythoner{}
	pro.Coding()

}
