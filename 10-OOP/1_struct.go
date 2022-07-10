package main

import "fmt"

//type  定义一个名字为 myint 的 int，在这代表为int 取别名
type myint int

//type 定义一个名字为Book 的struct 结构体
//啥是结构体呢？结构体就是一个融合了多种不同的数据类型的一个复杂的数据类型
type Book struct {
	title string
	auth  string
}

//结构体传方法，值传递
func changeBook(book Book) {
	book.auth = "li4"
}

//结构体传方法，引用传递
func changeBook2(book *Book) {
	book.auth = "li4"
}

func main() {

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhangsan"
	fmt.Printf("%v\n", book1)
	//传值
	changeBook(book1)
	fmt.Printf("%v\n", book1)
	//传地址
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
