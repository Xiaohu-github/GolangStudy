package main

import "fmt"

var ga int = 100
var gb = 200

// 四种变量的声明方式
func main() {
	// 方法一：声明一个整型变量，不赋值，默认值为 0
	var a int
	fmt.Println("a =", a)
	fmt.Printf("type a = %T\n", a)

	// 方法二：声明一个整型变量，初始化一个值
	var b int = 100
	fmt.Println("b =", b)
	fmt.Printf("type b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("type bb = %T , bb = %s\n", bb, bb)

	// 方法三：声明一个变量，不规定类型，默认会根据赋值来判断数据类型
	var c = 200
	fmt.Println("c =", c)
	fmt.Printf("type c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("type cc = %T , cc = %s\n", cc, cc)

	// 方法四（常用）：省去var 关键字，直接自动匹配
	// 注意：方法四 :=不能声明全局变量，只能在函数体内使用
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := 3.141231321335
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	g := "ABCD"
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	//==============
	fmt.Println("ga = ", ga, ", gb = ", gb)

	//多变量声明
	var xx, yy int = 100, 200
	fmt.Println("xx =", xx, ", yy = ", yy)

	var kk, ll = 200, "abcddd"
	fmt.Println("kk = ", kk, ", ll =", ll)

	var (
		vv int    = 500
		jj string = "dasdsda"
		tt bool   = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj, ", tt= ", tt)

}
