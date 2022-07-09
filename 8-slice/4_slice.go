/*
 len ,cap
 切片追加元素:append
*/
package main

import "fmt"

func main() {
	//定义一个切片，长度len=3，空间大小cap=5
	var numbers []int = make([]int, 3, 5)
	fmt.Printf("len = %d , cap = %d , slice = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers追加一个元素1，numbers 长度len=4，[0,0,0,1]， cap=5
	numbers = append(numbers, 1)
	fmt.Printf("len = %d , cap = %d , slice = %v\n", len(numbers), cap(numbers), numbers)

	//如果追加元素超过定义的cap，len > cap，那么切片会自动追加一倍cap空间 cap:5 -> cap:10
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	fmt.Printf("len = %d , cap = %d , slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("=====-- numbers2 --=====")
	//定义一个切片，长度len=3，空间大小默认cap=len=3
	numbers2 := make([]int, 3)
	fmt.Printf("len = %d , cap = %d , slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d , cap = %d , slice = %v\n", len(numbers2), cap(numbers2), numbers2)
}
