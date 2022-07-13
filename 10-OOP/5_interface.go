package main

import "fmt"

/**
interface{} 是万能数据类型：int\string\float32\float64\...都实现了 interface{}

如果你的方法不确定定传值类型 就把arg 定义为interface{}
*/
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//如何区分interface{} 此时引用的数据类型到底是什么
	//给 interface{} 提供“类型断言”的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg not string....")
	} else {
		fmt.Println("arg is string type, value=", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book1 := Book{"Golang"}
	myFunc(book1)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
