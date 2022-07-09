package main

import "fmt"

//如果方法传的是动态数组，它就不是值传递而是引用传递
//所以方法内修改的同时外部也会修改
func printArr(arr []int) {
	//_表示匿名的变量
	for _, value := range arr {
		fmt.Println("value=>", value)
	}
	arr[0], arr[1] = 100, 200
}

func main() {
	//slice:动态数组(切片)
	myarray := []int{1, 2, 3, 4}

	fmt.Printf("myArray types=> %T\n", myarray)

	printArr(myarray)

	fmt.Println("==============")
	for _, value := range myarray {
		fmt.Println("value=>", value)
	}
}
