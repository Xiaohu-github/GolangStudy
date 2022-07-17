/**
*有缓存的channle
*当channel缓存满了，再向里面传数据，就会阻塞
*当channel缓存为空，读取数据也会阻塞
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //有缓存的channel
	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子go程结束...")
		for i := 1; i <= 4; i++ {
			c <- i
			fmt.Println("子go程正在运行：元素=", i, "len(c) = ", len(c), ",cap(c) = ", cap(c))
		}
	}()

	//主程休眠2秒...此时子程应该已经把三个元素放到管道中了
	time.Sleep(5 * time.Second)

	for i := 1; i <= 4; i++ {
		num := <-c
		fmt.Println("num=", num)
	}

	fmt.Println("main 结束")
}
