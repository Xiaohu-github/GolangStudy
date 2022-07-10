/**
map的使用
删除：delete(map变量名，map键名)
增加：变量名[新键名]=新值
遍历：for k,v=range 变量名{
	操作...
}
*/
package main

import "fmt"

//map和slice一样是引用传递
func changeMap(myMap map[int]string) {
	myMap[5] = "C#"
	myMap[6] = "AA"
}

func main() {
	//声明
	myMap := make(map[int]string)

	//添加
	myMap[1] = "JAVA"
	myMap[3] = "PHP"
	myMap[2] = "C++"
	myMap[5] = "AAA"
	changeMap(myMap)
	fmt.Println("map==>", myMap)

	//删除
	delete(myMap, 6)
	myMap[4] = "python"

	copyMap := make(map[int]string)
	//遍历
	for key, value := range myMap {
		fmt.Println("key =", key)
		fmt.Println("value=", value)
		//拷贝
		copyMap[key] = value
	}
	fmt.Println("copyMap==>", copyMap)
}
