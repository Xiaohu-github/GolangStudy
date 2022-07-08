package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a, b := 10, 20
	// swap
	swap(&a, &b)
	println("a =", a, "b =", b)

	var p *int = &a

	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println("*p -> a =", *p)

	//二级指针
	var pp **int = &p

	fmt.Println(&p)
	fmt.Println(pp)
	fmt.Println("**pp -> *p -> a =", **pp)
}
