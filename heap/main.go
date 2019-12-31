package main

import (
	"fmt"
	"github.com/fyniny/order/data"
)

func main() {
	arr, o := data.RandomArrayInt(9)
	fmt.Println(arr)
	heap(arr)
	fmt.Println(arr)
	fmt.Println(o)
}
func heap(arr []int) {
	// 建立最大堆
	buildHeap(arr)
	l := len(arr)
	for pos := l; pos > 0; pos-- {
		// 选择堆顶元素到适合位置并减少堆元素数量
		arr[pos-1], arr[0] = arr[0], arr[pos-1]
		// 向下重新调整最大堆
		adjust(arr, 0, pos-1)
	}
}
// 设计原则：堆顶元素为0
func buildHeap(arr []int) {
	l := len(arr)
	for i := l / 2 - 1; i >= 0; i-- {
		adjust(arr, i, l)
	}
}
func adjust(arr []int, k, l int) {
	for i := 2 * (k + 1) - 1; i < l ; i = 2 * (i + 1) - 1 {
		if i < l - 1 && arr[i] < arr[i+1] {
			i++
		}
		if arr[k] >= arr[i] {break} else {
			arr[k], arr[i] = arr[i], arr[k]
			k = i
		}
	}
}
