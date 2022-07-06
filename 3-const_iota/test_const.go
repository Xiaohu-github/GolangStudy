package main

import "fmt"

//const 常用来来定义枚举
const (
	BEIJING   = 0
	SHANGHAI  = 1
	SHENGZHEN = 2
)

const (
	//可以在const() 添加一个字段 iota ,每行iota都会累加1，第一行iota默认值为0
	BEIJING_A   = iota //0
	SHANGHAI_A         //1
	SHENGZHEN_A        //2
)

const (
	//iota计算,每行iota都会累加1同时乘以100，第一行iota默认值为0
	BEIJING_B   = 100 * iota //iota = 0, 100*0
	SHANGHAI_B               //iota = 1, 100*1
	SHENGZHEN_B              //iota = 2, 100*2
)

const (
	//iota 是逐行累加1，同一行是相同的iota,每行逐一计算。
	//iota 只能出现在 const() 里面
	a, b = iota + 1, iota + 2 //iota = 0, a = iota+1, b = iota+2  | a=1,b=2
	c, d                      //iota = 1, c = iota+1, d = iota+2	| c=2,d=3
	e, f                      //iota = 2, e = iota+1, f = iota+2  | e=3,f=4

	g, h = iota * 2, iota * 3 //iota = 3, g = iota*2, h = iota*3	| g=6,h=9
	i, j                      //iota = 4, i = iota*2, j = ipta*3	| i=8,j=12
)

func main() {
	//常量（只读属性）
	const length int = 10

	fmt.Println("length =", length)

	fmt.Println("逐一定义枚举：BEIJING =", BEIJING, ", SHANGHAI =", SHANGHAI, ", SHENGZHEN =", SHENGZHEN)
	fmt.Println("iota累加定义枚举：BEIJING_A =", BEIJING_A, ", SHANGHAI_A =", SHANGHAI_A, ", SHENGZHEN_A =", SHENGZHEN_A)
	fmt.Println("iota计算枚举：BEIJING_B =", BEIJING_B, ", SHANGHAI_B =", SHANGHAI_B, ", SHENGZHEN_B =", SHENGZHEN_B)

	fmt.Println("a =", a, " b =", b)
	fmt.Println("c =", c, " d =", d)
	fmt.Println("e =", e, " f =", f)
	fmt.Println("g =", g, " h =", h)
	fmt.Println("i =", i, " j =", j)

}
