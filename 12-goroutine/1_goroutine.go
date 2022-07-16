/**
* 	执行一个函数
*		func()
*		开启一个协程执行这个函数
*		go func()
 */
package main

import (
	"fmt"
	"time"
)

func newTask() {
	//死循环
	i := 0
	for {
		i++
		fmt.Printf("------ new Goroutine : i == %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//创建go程（协程）去执行newTask，此时mian 和 newTask 是异步的，main 不会去等 newTask 执行完，就会结束
	go newTask()
	time.Sleep(1 * time.Second)

	//以下给main也套一个死循环，那么这两个方法就会并行执行，main相当于主协程，主程结束，从程也会结束
	i2 := 0
	for {
		i2++
		fmt.Printf("--- main Goroutine : i2 === %d\n", i2)
		time.Sleep(1 * time.Second)
	}
}
