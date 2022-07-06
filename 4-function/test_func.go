package main

import (
	"fmt"
)

func foo1(a string, b int) int {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := 100
	return c
}

//多返回值，匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("=== foo2 ===")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	return 66, 77
}

//多返回值，有形参名称
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("=== foo3 ===")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	//r1 r2 属于foo3的形参,默认初始值为0
	//r1 r2 作用域属于foo3整个函数体{}
	fmt.Println("r1 =", r1)
	fmt.Println("r2 =", r2)

	//给有名称形参赋值，返回时可省略返回参数
	r1 = 100
	r2 = 200

	return
}

//简写，当实参或者形参类型相同时，可简写
func foo4(a, b int) (r1, r2 int) {
	fmt.Println("=== foo4 ===")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	r1 = 1000
	r2 = 2000

	return
}

func main() {
	c := foo1("abc", 100)
	fmt.Println("c =", c)

	d, f := foo2("abc", 100)
	fmt.Println("d =", d)
	fmt.Println("f =", f)

	k, l := foo3("efg", 100)
	fmt.Println("k =", k)
	fmt.Println("l =", l)

	i, j := foo4(100, 200)
	fmt.Println("i =", i)
	fmt.Println("j =", j)
}
