package main

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
}
