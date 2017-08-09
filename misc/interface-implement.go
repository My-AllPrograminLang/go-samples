// Exploring what happens when we attempt to implement an interface "partially"
// (not all of the interface's methods), and the kind of errors the Go
// compiler emits.
package main

import "fmt"

type MyInterface interface {
	Foo() string
	Bar() string
}

type MyImplFull struct {
}

func (m *MyImplFull) Foo() string {
	return "Foo full"
}

func (m *MyImplFull) Bar() string {
	return "Bar full"
}

type MyImplPartial struct {
}

func (m *MyImplPartial) Foo() string {
	return "Foo partial"
}

func DoIt(obj MyInterface) {
	fmt.Println(obj.Foo())
}

func main() {
	f := &MyImplFull{}
	DoIt(f)

	// This will fail because MyImplPartial doesn't implement MyInterface fully,
	// so we cannot pass it to a function taking a MyInterface.
	//f2 := &MyImplPartial{}
	//DoIt(f2)
}
