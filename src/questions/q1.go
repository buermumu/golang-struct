package main

import "fmt"

/**
题目: 在一个二维数据中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增排序，请完成一个函数，输入这样一个整数，判断数组中是否包含该整数；
*/

var list [10][10]int

func findNumber(list [10][10]int, number int) {
	var line, cloum, total int
	line = 0
	cloum = 9
	total = 0
	for {
		total++
		num := list[line][cloum]
		if number == num {
			fmt.Printf("line:%d, cloum:%d, num:%d, number:%d, total:%d", line, cloum, num, number, total)
			break
		}
		if number > num {
			line++
		} else {
			cloum--
		}
		if line > 9 || cloum < 0 {
			fmt.Println("not find this number")
			break
		}
	}
}

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			list[i][j] = (i + 1) * (j + 1)
		}
	}
	findNumber(list, 54)
}
