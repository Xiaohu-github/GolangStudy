/*
*单流程下一个go只能监控一个channel状态，select可以完成多个channel状态的监控
*和switch 有点相似
	select{
			case <- chan1: 从chan1获取数据，就进行case处理语句
			case chan2<-1: 如果成功向chan2写入数据，就进行case处理语句
			default: 如果上面都没成功，就进入default处理语句
	}
*
*/
package main

import "fmt"

func fibonacli(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x: //将x发送给c管道，发送成功就执行下面语句 斐波那契 累加
			x = y
			y = x + y
		case <-quit: //如果quit管道里面有值，就代表循环结束了
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c) //由于c管道无缓存，所以循环拿c的值时，如果c没有值该channel会等待,直到有值给c的时候就会打印
		}
		quit <- 0
	}()

	fibonacli(c, quit)
}
