package main

import (
	"fmt"
	"reflect"
)

type Plane struct {
	color string
}
type Bird struct {
	Plane
}

func (plane Plane) fly() {
	fmt.Println(plane.color)
}

func (bird Bird) fly() {
	fmt.Println(bird.color + " change")
	bird.Plane.fly()
}

func add[T int | float64 | string](a, b T) T { return a + b }

func add2[T any](a, b T) T { return a }



type stu struct {
	name string
}

func testRefect() {
	var a = &stu{"hello"}

	var b = interface{}(a)

	var bType = reflect.TypeOf(b)
	fmt.Println(bType.Elem().Kind())
	fmt.Println(bType.Kind())
	var bValue = reflect.ValueOf(b)
	fmt.Println(bValue.Pointer())
	fmt.Println(bValue.Type())

}

func main() {
	// var p = Plane{"p color"}

	// var b = Bird{Plane: Plane{"b color"}}

	// p.fly()
	// b.fly()

	// var t int64 = 8
	// var a = add2(t, 3)
	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("a: %T\n", a)
	testRefect()

}

