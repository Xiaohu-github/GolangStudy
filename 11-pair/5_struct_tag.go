/**
*  结构体标签
*  标签的意义是在其他包导入使用时能够对属性进行说明，以此判断该属性如何用意义是什么
 */
package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"` //当前的属性绑定了两个标签,对属性进行说明
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	// t = t.Elem() //如果传过来的是结构体的地址得先获取元素，再通过元素调用Numfild
	t := reflect.TypeOf(str)
	fmt.Println("t TypeOf is", t)
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", taginfo, "doc:", tagdoc)
	}
}

func main() {
	// var re resume
	re := resume{}
	findTag(re)
}
