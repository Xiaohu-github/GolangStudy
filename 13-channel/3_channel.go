/**
*管道关闭
*channel 不像文件一样需要经常关闭，除非你确定该管道确实没啥数据了，或者因为某些原因才去关闭管道
关闭后，无法再向channel发送数据
关闭后，其他管道可以继续从channel拿数据，但此时肯定是没有数据的。
对于 空管道（nil channel）来说，无论收发都会被阻塞
*/
package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//通过close关闭channel
		close(c)
	}()

	for {
		//if 的简写
		//ok 如果为true表示channel没有关闭，如果为false表示channel 关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main Finish.")
}
