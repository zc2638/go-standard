package main

import "container/heap"

// heap包实现了对任意类型（实现了heap.Interface接口）的堆操作
func main() {

	// 初始化结构体
	h := &IntHeap{1, 4, 2, 7, 8, 9, 3, 6}
	// 初始化堆
	heap.Init(h)
	// 添加元素并重新排序
	heap.Push(h, 5)
	// 移除最小元素返回一个新的interface,根据sort排序规则
	heap.Pop(h)
	// 移除指定位置的interface，返回一个新的interface
	heap.Remove(h, 5)
	// 移除指定位置的interface，并修复索引
	heap.Fix(h, 5)
}

// 声明结构体
type IntHeap []int

// 创建sort.Interface接口的Len方法
func (h IntHeap) Len() int { return len(h) }

// 创建sort.Interface接口的Less方法
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// 创建sort.Interface接口的Swap方法
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// 创建heap.Interface接口的添加方法
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

// 创建heap.Interface接口的移除方法
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
