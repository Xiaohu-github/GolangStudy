/*
* 一个变量包含了 type 和 value：这俩组合统称为pair
* 而 type 又分为两类static type(int、string...)静态类型 和concrete type（interface所指向的具体的数据类型）实体类型
* 一个变量不管如何赋值给其他变量，它本身的 pair<type,value> 永远不变
 */
package main

import "fmt"

func main() {
	var a string
	//pair<type:tring,value:"accccd">
	a = "accccd"

	//pair<type:tring,value:"accccd">
	var allType interface{}
	allType = a

	//断言allType 是一个string 类型 （断言意思就是“我觉得” => 我觉得 allType 是一个 String 类型，然后系统返回 值 和 “你觉得”的结果）
	str, _ := allType.(string)
	fmt.Println(str)

}
