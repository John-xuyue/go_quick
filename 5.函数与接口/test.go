package main

import "fmt"



func soo() int {
	fmt.Println("enter soo")

	defer func() { //去掉这个defer试试，看看panic的流程。把这个defer放到soo函数末尾试试。把这个defer移到main()里试试。
		//recover必须在defer中才能生效
		if panic_value := recover(); panic_value != nil {
			fmt.Printf("soo函数中发生了panic:%v\n", panic_value)
		}
		fmt.Printf("test\n")
	}()

	fmt.Println("regist recover")

	defer fmt.Println("hello")
	defer func() {
		n := 0
		_ = 3 / n //除0异常，发生panic，下一行的defer没有注册成功
		defer fmt.Println("how are you")
	}()

	return 1
}


type Player interface {
	do(ball string) int
}


type Student struct {
	name string
}
func (s Student) do(name string) int {
	s.name = name
	return 1
}

type Employee struct {
	name string
}
func (e *Employee) do(name string) int {
	e.name = name
	return 2
}

func main() {
	
	// a := soo()
	// fmt.Printf("a: %v\n", a)

	var p Player
	var s = Student{"stu"}
	var e = Employee{"emp"}
	fmt.Printf("s: %v\n", s)
	fmt.Printf("e: %v\n", e)


	p = &s
	p.do("sss")
	fmt.Printf("s: %v\n", s)

	p = s
	p.do("sss")
	fmt.Printf("s: %v\n", s)

	p = &e
	p.do("eee")
	fmt.Printf("e: %v\n", e)
}
