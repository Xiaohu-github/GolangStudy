/*
类的继承
权限只有大小写，GO简化了权限构造
*/
package main

import "fmt"

type Human struct {
	Name, Sex string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (this *Human) walk() {
	fmt.Println("Human.walk()...")
}
func (this *Human) info() {
	fmt.Println("======INFO=====")
	fmt.Println("name=", this.Name)
	fmt.Println("Sex=", this.Sex)
}

type SuperMan struct {
	Human //继承父类，将类名放到子类结构体中
	Level int
}

func (this *SuperMan) fly() {
	fmt.Println("SuperMan.fly()...")
}

func main() {
	zhangs := Human{Name: "zhan3", Sex: "man"}
	zhangs.info()
	zhangs.Eat()
	zhangs.walk()

	lisi := SuperMan{Human{"li4", "woman"}, 100}
	lisi.info()
	lisi.fly()
	lisi.Eat()
	lisi.walk()

	var wang5 SuperMan
	wang5.Name = "wangwu"
	wang5.Sex = "woman"
	wang5.info()
	wang5.fly()
	lisi.Eat()
	lisi.walk()

}
