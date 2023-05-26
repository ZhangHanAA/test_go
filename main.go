package main

import "fmt"

func main() {
	// 快速排序测试
	arr := []int{1, 3, 2, 5, 4, 7, 6, 9, 8}
	arr = quickSort(arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}

// 快速排序
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	left := make([]int, 0)
	right := make([]int, 0)
	mid := make([]int, 0)
	mid = append(mid, pivot)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else if arr[i] > pivot {
			right = append(right, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	left, right = quickSort(left), quickSort(right)
	myArr := append(append(left, mid...), right...)
	return myArr
}
