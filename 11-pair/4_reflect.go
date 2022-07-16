package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("User func called....")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "HU", 24}
	user.Call()
	DofileAndMethod(user)
}

func DofileAndMethod(input interface{}) {
	//获取变量的类型
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType)

	//获取变量的值
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	//通过类型获取类型里面的字段
	//1、获取变量的Type,通过Type得到NumField（字段总数）,然后遍历
	for i := 0; i < inputType.NumField(); i++ {

		//2、通过type得到每个filed（字段类型）
		field := inputType.Field(i)

		//3、通过Value得到字段的Interface()=>对应的value值
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s ：%v = %v\n", field.Name, field.Type, value)
	}

	var num int = inputType.NumMethod()

	//通过type 获取里面的方法
	for i := 0; i < num; i++ {
		m := inputType.Method(i)
		fmt.Printf("%s:%v\n", m.Name, m.Type)
	}
}
