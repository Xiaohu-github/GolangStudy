/**
map 和 slice 差不多
区别在于slice里面是值类型统一的数组，
而map里面的值是hash 的key =>value 形式
*/
package main

import "fmt"

func main() {
	//声明方法一：声明myMap1是一个map类型 key是string，valu是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}
	//在使用map前需要给map分配空间,和slice一样如果超出空间大小，go会自动扩充一倍
	myMap1 = make(map[string]string, 10)
	//由于map的存储为hash，所以它是无序的
	myMap1["one"] = "Java"
	myMap1["two"] = "c++"
	myMap1["three"] = "php"

	fmt.Println("myMap1 = ", myMap1)

	//声明方法二：采用:=方式声明，可以不必声明大小，当赋值时会自动扩充空间
	myMap2 := make(map[int]string)
	myMap2[1] = "JAVA"
	myMap2[3] = "PHP"
	myMap2[2] = "C++"

	fmt.Println("myMap2 = ", myMap2)

	//申明方法三：直接声明并赋值
	myMap3 := map[string]string{
		"one":   "aaa",
		"two":   "bbb",
		"three": "ccc",
	}
	fmt.Println("myMap3 = ", myMap3)

}
