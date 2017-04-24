package isort

import (
	_ "fmt"
)

type Sort struct {
	data  []int
	total int
}

func Init(data []int) *Sort {
	return &Sort{data: data}
}

/**
 * 直接插入排序
 */
func (s *Sort) Insert() *Sort {
	var tmp int
	length := s.Length()
	for i := 1; i < length; i++ {
		tmp = s.data[i]
		for j := i; j > 0; j-- {
			s.total++
			if tmp < s.data[j-1] {
				s.data[j], s.data[j-1] = s.data[j-1], tmp
			} else {
				break
			}
		}
	}
	return s
}

/**
 * 希尔排序
 */
func (s *Sort) Shell() *Sort {
	return s
}

/**
 * 快速排充
 */
func (s *Sort) Quick() *Sort {
	return s
}

/**
 * 超泡排序
 */
func (s *Sort) Bubble() *Sort {
	return s
}

/**
 * 选择排序
 */
func (s *Sort) Select() *Sort {
	return s
}

/**
 * 堆排序
 */
func (s *Sort) Heap() *Sort {
	return s
}

/**
 * 归并排序
 */
func (s *Sort) Merge() *Sort {
	return s
}

/**
 * 桶排序/基数排序
 */
func (s *Sort) Radix() *Sort {
	return s
}

func (s *Sort) Length() int {
	return len(s.data)
}
