package main

import (
	"github.com/fyniny/order/data"
)

func main() {
	arr, o := data.RandomArrayInt(10000)
	radix(arr, 10)
	data.Equal(arr, o)
}

type queue struct {
	data []int
	cap int
	head int
	tail int
	status int
	len int
}
func (q *queue) push(elem int) {
	if q.status != full {
		q.data[q.head] = elem
		q.head = (q.head + 1) % q.cap
		q.len++
		if q.head == q.tail {
			q.status = full
		} else if q.status == empty {
			q.status = work
		}
		return
	}
	panic("queue is full")
}
func (q *queue) front() int {
	if q.status != empty {
		return q.data[q.tail]
	}
	panic("queue is empty")
}
func (q *queue) pop() int {
	if q.status != empty {
		ret := q.data[q.tail]
		q.len--
		q.tail = (q.tail + 1) % q.cap
		if q.tail == q.head {
			q.status = empty
		} else if q.status == full {
			q.status = work
		}
		return ret
	}
	panic("queue is empty")
}
func (q *queue) size() int {
	return q.len
}
func (q *queue) init(size int) {
	q.cap = size
	q.data = make([]int, q.cap)
	q.head = 0
	q.tail = 0
	q.status = empty
	q.len = 0
}
const (
	empty = -1
	full = 1
	work = 0
)

func radix(arr []int, base int) {
	queues := make([]queue, base)
	l := len(arr)
	for i := 0; i < base; i++ {queues[i].init(l)}
	// 获取需要分配次数
	d := times(arr, base)
	div := 1
	for i := 0; i < d; i++ {
		// distribution
		for j := 0; j < l; j++ {
			index := arr[j] / div % base
			queues[index].push(arr[j])
		}
		// collection
		for j, k := 0, 0; j < len(queues); j++ {
			for 0 != queues[j].size() {
				arr[k] = queues[j].pop()
				k++
			}
		}
		div *= base
	}
}

func times(arr []int, base int) int {
	m := arr[0]
	for i := range arr {
		if arr[i] > m {
			m = arr[i]
		}
	}

	i := 1
	for ;m / base != 0; m /= base {
		i++
	}
	return i
}