package main

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var a = []int{1, 8, 2, 9, 0, 3, 3, 0, 9, 1}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

//思想：选取一个靶点，使得在一次循环遍历中左边的元素都小于靶点，右边的元素都大于靶点，再递归遍历左边和右边
func QuickSort(data []int, start, end int) {
	if start >= end {
		return
	}
	var i, j = start, end
	var mid = data[start]
	for {
		for data[j] >= mid && i < j {
			j--
		}
		for data[i] <= mid && i < j {
			i++
		}
		//每次比较总会出现的
		if i == j {
			break
		}
		data[i], data[j] = data[j], data[i]
	}
	//和靶点交换
	data[start], data[i] = data[i], data[start]

	QuickSort(data, start, i-1)
	QuickSort(data, i+1, end)
}
