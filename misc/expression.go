package main

import (
	"fmt"
	"strconv"
)

type Expression interface {
	Eval() float64
}

type Stringable interface {
	ToString() string
}

type Constant struct {
	value float64
}

func (c *Constant) Eval() float64 {
	return c.value
}

func (c *Constant) ToString() string {
  return strconv.FormatFloat(c.value, 'E', -1, 64)
}

type BinPlus struct {
	left Expression
	right Expression
}

func (bp *BinPlus) Eval() float64 {
	return bp.left.Eval() + bp.right.Eval()
}

func (bp *BinPlus) ToString() string {
	// The moment of truth is here... bp.left is an Expression, which does not
	// have a ToString method. Obviously this will only work if left and right
	// implement the Stringable interface. The type assertion makes this
	// expectation explicit and will panic otherwise.
	ls := bp.left.(Stringable)
	rs := bp.right.(Stringable)
	return fmt.Sprintf("(%s + %s)", ls.ToString(), rs.ToString())
}

func main() {
	fmt.Println("booya")

	// constants
	c := Constant{value: 26.4}

	fmt.Printf("c Eval = %g\n", c.Eval())
	fmt.Printf("c ToString = %s\n", c.ToString())

	c11 := Constant{value: 1.1}
	c22 := Constant{value: 2.2}
	c33 := Constant{value: 3.3}
	bp := BinPlus{left: &BinPlus{left: &c11, right: &c22}, right: &c33}

	fmt.Printf("bp Eval = %g\n", bp.Eval())
	fmt.Printf("bp ToString = %s\n", bp.ToString())
}
