package disign

import "math"

/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。
 */

type MinStack struct {
	index int
	arr []int
	arrMin []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{
		-1,
		[]int{},
		[]int{},
	}
	return stack
}


func (this *MinStack) Push(x int)  {
	if len(this.arr) > this.index + 1 {
		this.arr[this.index + 1] = x
		this.arrMin[this.index + 1] = min(x, this.GetMin())
		this.index++
	} else {
		this.arr = append(this.arr, x)
		this.arrMin = append(this.arrMin, min(x, this.GetMin()))
		this.index++
	}
}


func (this *MinStack) Pop()  {
	this.index--
}


func (this *MinStack) Top() int {
	return this.arr[this.index]
}


func (this *MinStack) GetMin() int {
	if this.index == -1 {
		return math.MaxInt32
	}
	return this.arrMin[this.index]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

