package main

import (
	"fmt"
)

func main() {
	//定义一个channel(管道)
	c := make(chan int)

	//go goroutine 一个 匿名函数
	go func() {
		defer fmt.Println("Goroutine 结束...")

		fmt.Println("Goroutine 运行中")

		c <- 666 //将666发送给c管道
	}()

	num := <-c //从c中接收数据

	//这块牵扯到一个执行先后的问题，由于main需要从管道中获取sub协程发过来的666，
	//所以无论哪个先跑完，只要main还没拿到666，channel都会进行阻塞等待，等待值发过来再进行下面的操作
	//变相进行了同步

	fmt.Println("num=", num)
	fmt.Println("main goroutine 结束...")
}
