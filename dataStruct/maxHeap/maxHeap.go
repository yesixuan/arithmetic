package maxHeap

import (
	"fmt"
	"math"
)

type MaxHeap struct {
	list []int
}

// 获取父亲位置
func parentIndex(index int) int {
	if index <= 0 {
		return -1
	}
	return int(math.Floor(float64((index - 1) / 2)))
}

// 获取左孩子位置
func leftIndex(index int) int {
	return index*2 + 1
}

// 获取有孩子位置
func rightIndex(index int) int {
	return index*2 + 1
}

// 交换两个元素位置
func (maxHeap *MaxHeap) swap(a int, b int) {
	tmp := maxHeap.list[a]
	maxHeap.list[a] = maxHeap.list[b]
	maxHeap.list[b] = tmp
}

func (maxHeap *MaxHeap) siftup(index int) {
	list := maxHeap.list
	for index > 0 && list[index] > list[parentIndex(index)] {
		maxHeap.swap(index, parentIndex(index))
		index = parentIndex(index)
	}
}

// 添加元素
func (maxHeap *MaxHeap) Add(e int) {
	maxHeap.list = append(maxHeap.list, e)
	maxHeap.siftup(len(maxHeap.list) - 1)
}

// TODO 我来看看
func (maxHeap *MaxHeap) Hepify(list []int) {
	maxHeap.list = list
	for i := parentIndex(len(maxHeap.list) - 1); i > 0; i-- {
		// i 位置进行下潜
	}
}

func (maxHeap *MaxHeap) Print() {
	fmt.Println(maxHeap.list)
}
