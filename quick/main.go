package main

import (
	"fmt"
	"github.com/fyniny/order/data"
)

func main() {
	arr, o := data.RandomArrayInt(25)
	quick(arr, 0, len(arr)-1)
	fmt.Println(arr)
	fmt.Println(o)
}

func quick(arr []int, low, high int) {
	if low < high {
		pos := partition(arr, low, high)
		quick(arr, low, pos-1)
		quick(arr, pos+1, high)
	}
}
func partition(arr []int, low, high int) (pos int) {
	// pivot 设定为左边第一个，故从最右边开始查找
	pivot := arr[low]
	for low < high {
		// 寻找右边小于pivot的序列下标
		for low < high && arr[high] >= pivot {high--}
		// 此时arr[high] < pilot， 所以将其设置给arr[low]
		arr[low] = arr[high]
		for low < high && arr[low] <= pivot {low++}
		arr[high] = arr[low]
	}
	// 退出循环时: low == high -> true
	arr[low] = pivot
	pos = low
	return
}