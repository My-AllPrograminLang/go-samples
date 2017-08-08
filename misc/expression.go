package main

import (
	"fmt"
	"strconv"
)

type Expression interface {
	Evaluate() float64
}

type Stringable interface {
	ToString() string
}

type Constant struct {
	value float64
}

func (c *Constant) Evaluate() float64 {
	return c.value
}

func (c *Constant) ToString() string {
  return strconv.FormatFloat(c.value, 'E', -1, 64)
}

type BinPlus struct {
	left Expression
	right Expression
}

func (bp *BinPlus) Evaluate() float64 {
	return bp.left.Evaluate() + bp.right.Evaluate()
}

func (bp *BinPlus) ToString() string {
	ls := bp.left.(Stringable)
	rs := bp.right.(Stringable)
	return fmt.Sprintf("(%s + %s)", ls.ToString(), rs.ToString())
}

func main() {
	fmt.Println("booya")

	// constants
	c := Constant{value: 26.4}

	fmt.Printf("c Evaluate = %g\n", c.Evaluate())
	fmt.Printf("c ToString = %s\n", c.ToString())

	c11 := Constant{value: 1.1}
	c22 := Constant{value: 2.2}
	c33 := Constant{value: 3.3}
	bp := BinPlus{left: &BinPlus{left: &c11, right: &c22}, right: &c33}

	fmt.Printf("bp Evaluate = %g\n", bp.Evaluate())
	fmt.Printf("bp ToString = %s\n", bp.ToString())
}
