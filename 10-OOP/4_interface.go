/*
多态是指在面向对象中能够根据使用类的上下文来重新定义或改变类的性质和行为
多态基本要素：
1、存在一个父类（接口）
2、子类实现父类全部方法
3、父类类型的变量指向子类的具体变量
*/
package main

import "fmt"

//本质是一个指针
//定义一个Animal 接口
type AnimalIf interface {
	Sleep()
	GetColor() string //颜色
	GetType() string  //获取动物类型
}

//定义一个类
type Cat struct {
	coler string //猫颜色
}

func (this Cat) Sleep() {
	fmt.Println("cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.coler
}

func (this *Cat) GetType() string {
	return "type:Cat..."
}

//定义一个类
type Dog struct {
	coler string //狗的颜色
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.coler
}

func (this *Dog) GetType() string {
	return "type:Dog..."
}

func showAnimal(animal AnimalIf) {
	animal.Sleep()
	fmt.Println("coler=", animal.GetColor())
	fmt.Println("kind=", animal.GetType())
}

func main() {
	cat := &Cat{"yellow"}
	showAnimal(cat)

	var animal AnimalIf
	animal = &Dog{"yellow"}
	animal.Sleep()
}
