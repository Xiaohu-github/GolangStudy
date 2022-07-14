/**
*reflect 反射
*TypeOf:获取变量的类型
*ValueOf:获取变量的值
 */
package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type :", reflect.TypeOf(arg))
	fmt.Println("value :", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345
	reflectNum(num)
}
