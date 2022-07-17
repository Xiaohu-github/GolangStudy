package main

import (
	"fmt"
	"time"
)

func main() {
	//go 一个 匿名函数，形参为空返回值为空
	// go func() {
	// 	defer fmt.Println("A.defer")

	// 	//匿名函数
	// 	func() {
	// 		defer fmt.Println("B.defer")
	// 		//退出当前Goroutine(go协程)
	// 		runtime.Goexit()
	// 		fmt.Println("B")
	// 	}()

	// 	fmt.Println("A")
	// }()

	//go协程 和主程是并行的 所以go协程并不会把里面的返回值返回给上级， “ v:= ” 获取不了
	//主程要想和go程通信，需要借助管道（channel）来使两个协程通信
	go func(a int, b int) bool {
		fmt.Println("a=", a, "b=", b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}

}
