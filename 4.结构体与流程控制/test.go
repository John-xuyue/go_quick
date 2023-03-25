package main

import "fmt"

func main() {

	var a int = 5

	switch a {
	case 5:
		{
			fmt.Printf("a: %v\n", a)
			fmt.Printf("a: %v\n", a + 1)
		}
	case 2:
		fmt.Printf("a: %v\n", a)
	}
}
