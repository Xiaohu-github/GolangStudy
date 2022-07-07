package lib1

import "fmt"

func init() {
	fmt.Println("lib1.init()...")
}

//当前lib1下的test方法
func Lib1Test() {
	fmt.Println("lib1Test()...")
}
