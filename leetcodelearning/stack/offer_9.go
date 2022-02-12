package stack

// 用两个栈实现队列

type CQueue struct {
	stack1 []int // 主栈
	stack2 []int // 副栈
}

func Constructor() CQueue {
	return CQueue{stack1: []int{}, stack2: []int{}}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack2) == 0 {
		if len(this.stack1) == 0 {
			return -1
		}
		// 逆序传给stack2
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		this.stack1 = []int{}
	}
	result := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1] // 去掉stack2栈顶元素
	return result
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
