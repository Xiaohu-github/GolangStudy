package main

import "fmt"

func deferfunc1() {
	fmt.Println("defer1")
}

func deferfunc2() {
	fmt.Println("defer2")
}

func deferfunc3() {
	fmt.Println("defer3")
}

func returnfunc() int {
	fmt.Println("return...")
	return 0
}

func returnAndDefer() int {
	//defer会在方法结束之前触发
	//多个defer按照压栈顺序执行：后进先出 3|2|1
	defer deferfunc1()
	defer deferfunc2()
	defer deferfunc3()

	fmt.Println("main::hello")
	fmt.Println("main::hello2")

	//return 和 defer 优先级先return 再 defer
	return returnfunc()
}

func main() {
	returnAndDefer()
}
