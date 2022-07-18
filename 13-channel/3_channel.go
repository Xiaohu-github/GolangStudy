/**
*管道关闭
*channel 不像文件一样需要经常关闭，除非你确定该管道确实没啥数据了，或者因为某些原因才去关闭管道
关闭后，无法再向channel发送数据
关闭后，其他管道可以继续从channel拿数据，缓存拿完后，此时管道就真的关闭了。
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

	// for {
	// 	//if 的简写
	// 	//ok 如果为true表示channel没有关闭，如果为false表示channel 关闭
	// 	if data, ok := <-c; ok {
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }

	//上面简写：可以使用range代替for循环拿数据
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main Finish.")
}
