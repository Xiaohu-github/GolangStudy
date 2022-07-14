/*
* pair 的例子
 */
package main

import (
	"fmt"
)

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriterBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book！")
}

func (this *Book) WriterBook() {
	fmt.Println("Writer a book")
}

func main() {
	//b:pair<type:Book,value:book{}地址>
	// var b *Book = &Book{}
	b := &Book{}

	fmt.Printf("b type is %T\n", b)
	b.ReadBook()

	//定义了一个 Reader interface r
	var r Reader //r：pair<type:,value:>

	r = b //r:pair<type:Book,value:book{}地址>
	fmt.Printf("r type is %T\n", r)
	r.ReadBook()

	//定义了一个 Reader interface w
	var w Writer //w：pair<type:,value:>

	w = r.(Writer) //w:pair<type:Boo k,value:book{}地址>

	w.WriterBook()
}
