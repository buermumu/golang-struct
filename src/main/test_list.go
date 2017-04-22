package main

import "list"
import "fmt"

func main() {
	L := list.New()
	L.InsertHead("a")
	L.InsertHead("b")
	L.InsertHead("c")
	L.InsertHead("d")
	L.InsertHead("e")
	L.Each()

	fmt.Println("---------------------------")

	ele := L.GetEleByIdx(1)
	fmt.Printf("ele.value :%s\n", ele.Value())
	L.InsertAfter(ele, "this is test")

	fmt.Println("---------------------------")

	L.Each()

	fmt.Printf("Len:%d \n", L.Len())

	L.Delete(ele.Next())
	fmt.Println("---------------------------")

	L.Each()
	fmt.Printf("Len:%d \n", L.Len())

}
