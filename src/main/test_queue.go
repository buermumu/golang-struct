package main

import (
	"fmt"
	"stack"
)

func main() {
	S1 := stack.New()
	S2 := stack.New()

	S1.Push("a")
	S1.Push("b")
	S1.Push("c")
	S1.Push("d")
	S1.Push("e")

	e := &stack.Element{}
	length := S1.Length()
	for i := 0; i < length; i++ {
		e = S1.Pop()
		S2.Push(e.Value())
	}

	length = S2.Length()
	for i := 0; i < length; i++ {
		fmt.Println(S2.Pop())
	}

	fmt.Println(S1, S2)
}
