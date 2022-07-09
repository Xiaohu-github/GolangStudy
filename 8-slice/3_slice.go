package main

import "fmt"

func printSlice(sli []int) {
	fmt.Printf("len = %d , slice = %v\n", len(sli), sli)
}

//slice 几种声明方式
func main() {
	//方法一：声明slice1是一个切片，并且初始化，默认值：1，2，3。 长度len为3
	slice1 := []int{1, 2, 3}
	printSlice(slice1)

	//方法二：声明slice2是一个切片，仅声明，无法直接使用如：slice2[0] = 100
	var slice2 []int
	//判断是否为空切片
	if slice2 == nil {
		fmt.Println("slice2是一个空切片")
	}
	//给slice2初始化，分配5个空间，初始化值为0
	slice2 = make([]int, 5)
	slice2[0] = 100
	printSlice(slice2)

	//方法三：声明slice3是一个切片，同时分配空间，3个空间，初始化为0
	// var slice3 []int = make([]int, 3)
	slice3 := make([]int, 3) //简写，通过:= 推导出slice3是一个切片
	printSlice(slice3)

}
