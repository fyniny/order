package main

import (
	"fmt"
	"github.com/fyniny/order/data"
)

func main() {
	arr, o := data.RandomArrayInt(4)
	insertOrder(arr)
	fmt.Println(arr)
	fmt.Println(o)
}

func insertOrder(arr []int) {
	var (
		l = len(arr)
		j = 0
	)
	for i := 1; i < l; i++ {
		tmp := arr[i]
		for j = i-1; j >= 0; j-- {
			if arr[j] <= tmp {
				break
			}
			arr[j+1] = arr[j]
		}
		arr[j+1] = tmp
	}
}
