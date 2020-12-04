package main

import (
	"container/heap"
	"fmt"
)

type priorityQueue []int

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	p := x.(int)
	*pq = append(*pq, p)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	p := old[n-1]
	*pq = old[0 : n-1]
	return p
}

func main() {
	pq := priorityQueue{1, 8, 8, 7, 2, 2, 5, 3, 5, 3, 1}
	heap.Init(&pq) //集合转变成最小堆
	for _, v := range pq {
		fmt.Print(v)
	}
	fmt.Println()
	pq[5] = -100 //在修改第i个元素后，调用本函数修复堆，比remove第i个元素后push新元素更有效率。
	heap.Fix(&pq, 5)
	for _, v := range pq {
		fmt.Print(v)
	}
	fmt.Println()
	heap.Remove(&pq, 5)
	heap.Push(&pq, -10)
	for _, v := range pq {
		fmt.Print(v)
	}
	fmt.Println()
	heap.Push(&pq, -5) //插入一个元素到最小堆，调整
	for _, v := range pq {
		fmt.Print(v)
	}
	fmt.Println()
	fmt.Println(heap.Pop(&pq)) //删除返回堆顶，即最小值，调整
	for _, v := range pq {
		fmt.Print(v)
	}
}
