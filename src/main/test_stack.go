package main

import "stack"
import "fmt"

func main() {
	S := stack.New()
	S.Push(nil)
	fmt.Println(S)
	p := S.Pop()
	t := S.Clear()
	fmt.Println(S, t, p)
}
