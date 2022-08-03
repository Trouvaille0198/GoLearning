package stack

type MinStack struct {
	stack, minStack []int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.minStack) == 0 || this.minStack[len(this.minStack)-1] >= val {
		this.minStack = append(this.minStack, val)
	}

}

func (this *MinStack) Pop() {
	topVal := this.Top()
	this.stack = this.stack[:len(this.stack)-1]
	if this.minStack[len(this.minStack)-1] == topVal {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
