package main

import "fmt"

type calc interface {
	add() int
	mul() int
}

type variable struct {
	a int
	b int
}

func (c variable) add() int {
	d := c.a + c.b
	return d
}

func (c variable) mul() int {
	d := c.a * c.b
	return d
}

func math(c calc) {
	fmt.Println(c.add())
	fmt.Println(c.mul())
}

func main() {
	t := variable{a: 5, b: 10}
	math(t)
}
