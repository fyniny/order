package main

import (
	"fmt"
	"github.com/fyniny/order/data"
)

func main() {
	arr, o := data.RandomArrayInt(7)
	shellOrder(arr)
	fmt.Println(arr)
	fmt.Println(o)
}

func shellOrder(arr []int) {
	gap := len(arr) / 2
	for ; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			j := i - gap
			tmp := arr[i]
			for ; j >= 0; j -= gap {
				if tmp > arr[j] {
					break
				}
				arr[j + gap] = arr[j]
			}
			arr[j + gap] = tmp
		}
	}
}
