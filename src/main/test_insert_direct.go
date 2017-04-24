package main

import (
	"fmt"
	"isort"
	_ "math/rand"
)

func main() {
	data := []int{17, 11, 19, 1000, 18, 0, -1, -32, -2, 12, 3, 6, 4, 84, 19, 199}
	//data := []int{1, 2, 3, 4, 5, 6, 7}
	//var data []int = make([]int, 1000)
	//for i := 0; i < 1000; i++ {
	//	data[i] = rand.Intn(1000000)
	//}
	s := isort.Init(data)
	r := s.Insert()
	fmt.Println(r)
}
