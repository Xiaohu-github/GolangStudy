/**
定义一个类
首先声明一个结构体(类)==>注：如果类名首字母大写表示共有(public),其他包能访问
type 结构体名 struct {
	结构体属性==>注：如果属性大写表示该属性能对外访问（public），否则只能内部访问
}
然后将方法绑定结构体
func (this 结构体名[类名]) 方法名 返回值类型（有就写没有就算了）{
	方法
}
大写：外部包可用
小写：仅内部包可用
*/
package main

import "fmt"

//=====类=====
//定义一个hero结构体
//go 语言中的类其实就是结构体绑定方法和属性
type Hero struct {
	Name      string
	Ad, Level int
}

//给结构体定义方法 this 是当前对象，谁调的我对象就是谁
func (this Hero) show() {
	fmt.Println("Hero-name:", this.Name)
	fmt.Println("Hero-Ad:", this.Ad)
	fmt.Println("Hero-Level:", this.Level)
}

func (this Hero) getName() string {
	return this.Name
}

//如果要修改原对象的属性，this 得是对象的指针才行，否则该 this 是拷贝了一个对象，值更改不会发生变化
func (this *Hero) setName(newName string) {
	this.Name = newName
}

//该 this 是拷贝对象,更改值只会在该方法中生效，值更改不会改变原对象的属性
func (this Hero) setName2(newName string) {
	this.Name = newName
	fmt.Println("副本this->Hero的name:", this.Name)
}

//=====以上组成一个类=====

func main() {
	//创建一个对象,当前对象是类的实例化
	hero := Hero{Name: "张三", Ad: 100, Level: 1}

	hero.show()
	hero.setName2("lisi")
	hero.setName("FUCk")
	hero.show()

}
