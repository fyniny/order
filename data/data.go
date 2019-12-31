package data

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomArrayInt(l int) (order, disOrder []int) {
	arr := make([]int, l)
	rand.Seed(time.Now().Unix())
	for j := 0; j < l; j++ {
		arr[j] = rand.Int() % (2 * l)
	}
	order = arr
	disOrder = insertOrder(arr)
	return
}

func insertOrder(arr []int) []int {
	ret := make([]int, 1, 1 + len(arr))
	ret = append(ret, arr...)
	for i := 2; i < 1 + len(arr); i++ {
		if ret[i] < ret[i-1] {
			j := i - 1
			ret[0] = ret[i]
			for ; ret[0] < ret[j]; j-- {
				ret[j+1] = ret[j]
			}
			ret[j+1] = ret[0]
		}
	}
	return ret[1:]
}

func Equal(arr1, arr2 []int) {
	if len(arr1) != len(arr2) {
	goto ERR
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			goto ERR
		}
	}
	fmt.Println("ok")
	return
ERR:
	fmt.Println("failed")
	fmt.Println(arr1)
	fmt.Println(arr2)
	return
}