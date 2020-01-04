package main

import (
	"fmt"
	"github.com/fyniny/order/data"
)

func main() {
	l := 20
	arr, o := data.RandomArrayInt(l)
	f := merge(l)
	mergeSort(arr, 0, len(arr)-1, f)
	fmt.Println(arr)
	fmt.Println(o)
}

func merge(l int) func(arr []int, low, mid, high int) {
	// 辅助数组
	helper := make([]int, l)
	return func(arr []int, low, mid, high int) {
		i, j, k := 0, 0, 0
		for k := low; k <= high; k++ {
			helper[k] = arr[k]
		}
		for i, j, k = low, mid+1, low; i <= mid && j <= high; k++ {
			if helper[i] > helper[j] {
				arr[k] = helper[j]
				j++
			} else {
				arr[k] = helper[i]
				i++
			}
		}
		for ;i <= mid; k++ {
			arr[k] = helper[i]
			i++
		}
		for ;j <= high; k++{
			arr[k] = helper[j]
			j++
		}
	}
}
func mergeSort(arr []int, low, high int, mergeFunc func([]int, int, int, int)) {
	if low < high {
		mid := (low + high) / 2
		mergeSort(arr, low, mid, mergeFunc)
		mergeSort(arr, mid+1, high, mergeFunc)
		mergeFunc(arr, low, mid, high)
	}
}
