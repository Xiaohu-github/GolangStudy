package main

import (
	"fmt"
)

//定长数组如果通过方法传值，
//那么传值的类型必须相同,定义是[10]int就必须传[10]int类型，
//且数据为值传递
func printArray(myarray [10]int) {

	fmt.Printf("myArray types=> %T\n", myarray)
	//range 会根据数据类型返回
	for index, value := range myarray {
		fmt.Println("index=>", index, ", value=>", value)
	}

}

func main() {
	//定长数组
	var myArray [10]int
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [10]int{1, 2, 3}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	printArray(myArray2)
	printArray(myArray3)

}
