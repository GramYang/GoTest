package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	heap1()
	//list1()
	//ring1()
}

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

func heap1() {
	pq := priorityQueue{1, 8, 8, 7, 2, 2, 5, 3, 5, 3, 1}
	heap.Init(&pq) //集合转变成最小堆
	for _, v := range pq {
		fmt.Print(v)
	} //11232857538
	fmt.Println()
	pq[5] = -100 //在修改第i个元素后，调用本函数修复堆，比remove第i个元素后push新元素更有效率。
	heap.Fix(&pq, 5)
	for _, v := range pq {
		fmt.Print(v)
	} //-1001132257538
	fmt.Println()
	heap.Remove(&pq, 5)
	heap.Push(&pq, -10)
	for _, v := range pq {
		fmt.Print(v)
	} //-100-10131857532
	fmt.Println()
	heap.Push(&pq, -5) //插入一个元素到最小堆，调整
	for _, v := range pq {
		fmt.Print(v)
	} //-100-10-5311575328
	fmt.Println()
	fmt.Println(heap.Pop(&pq)) //-100，删除返回堆顶，即最小值，调整
	for _, v := range pq {
		fmt.Print(v)
	} //-101-532157538
}

type element struct {
	age  int
	name string
}

func test1(es ...element) {
	l := list.New()
	for k := range es {
		l.PushBack(es[k])
	}
	for e := l.Front(); e != nil; e = e.Next() {
		elem, ok := e.Value.(element)
		if ok {
			fmt.Printf("Age: %d\n", elem.age)
			fmt.Printf("Name: %s\n", elem.name)
		}
	}
}

func list1() {
	e1 := element{24, "田所浩二"}
	e2 := element{20, "蔡徐坤"}
	e3 := element{26, "孙笑川"}
	test1(e1, e2, e3)
}

func ring1() {
	r := ring.New(10)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	for i := 0; i < r.Len(); i++ {
		fmt.Print(r.Value)
		r = r.Next()
	} //0123456789
	fmt.Println()
	r = r.Move(6)
	fmt.Println(r.Value) //6
	for i := 0; i < r.Len(); i++ {
		fmt.Print(r.Value)
		r = r.Next()
	}
	fmt.Println() //6789012345
	r1 := r.Unlink(19)
	for i := 0; i < r1.Len(); i++ {
		fmt.Print(r1.Value)
		r = r1.Next()
	}
	fmt.Println() //777777777
}
