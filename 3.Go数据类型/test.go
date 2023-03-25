package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	var b = append(a, 7)
	var c = a[:2]

	fmt.Printf("&a: %p\n", &a)
	fmt.Printf("a[0]: %p\n", &a[0])
	fmt.Printf("a: %p\n", a)
	fmt.Printf("b: %p\n", &b)
	fmt.Printf("c: %p\n", &c)
}