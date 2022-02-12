package stack

// 包含min函数的栈
// 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
// 调用 min、push 及 pop 的时间复杂度都是 O(1)。

import (
	"container/list"
)

type MinStack struct {
	stack    *list.List // 主栈
	minStack *list.List // 记录最小元素的栈
}

/** initialize your data structure here. */

func Constructor() MinStack {
	oriList := list.New()
	oriMinList := list.New()
	return MinStack{stack: oriList, minStack: oriMinList}
}

func (this *MinStack) Push(x int) {
	this.stack.PushBack(x)
	if this.minStack.Len() == 0 || x <= this.minStack.Back().Value.(int) {
		this.minStack.PushBack(x)
	}
}

func (this *MinStack) Pop() {
	deletedVal := this.stack.Remove(this.stack.Back())
	if deletedVal == this.minStack.Back().Value.(int) {
		this.minStack.Remove(this.minStack.Back())
	}
}

func (this *MinStack) Top() int {
	return this.stack.Back().Value.(int)
}

func (this *MinStack) Min() int {
	return this.minStack.Back().Value.(int)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
