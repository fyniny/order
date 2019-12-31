package main

import (
	"fmt"
	"github.com/fyniny/order/data"
)

func main() {
	arr, o := data.RandomArrayInt(100)
	selectOrder(arr)
	fmt.Println(arr)
	fmt.Println(o)
}

func selectOrder(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
