package main

import (
	"fmt"
	"github.com/fyniny/order/data"
)

func main() {
	arr,o := data.RandomArrayInt(20)
	bubble(arr)
	fmt.Println(arr)
	fmt.Println(o)
}

func bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
