package main

import (
	"fmt"
	"time"
)

func Add(a, b int) int {
	fmt.Println("add")
	return a + b
}

func F() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover")
		} else {
			fmt.Println("test")
		}
	}()
	fmt.Println("enter f")
	// panic(4)
}

func main() {
	// fmt.Println(runtime.NumCPU())
	// go Add(2, 4)

	// wg := sync.WaitGroup{}
	// wg.Add(10)

	// for i := 0; i < 10; i++ {
	// 	go func(a, b int) {
	// 		defer wg.Done()
	// 		fmt.Println(a + b)
	// 	}(i, i+1)
	// }
	// wg.Wait()

	go F()
	time.Sleep(2 * time.Second)
	fmt.Println("main")

}
