type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0{
		this.minStack = append(this.minStack, 0)
	} else if this.stack[this.minStack[len(this.minStack)-1]] > val{
		this.minStack = append(this.minStack, len(this.stack)-1)
	}
	// fmt.Println(this.stack, this.minStack)
}

func (this *MinStack) Pop() {
	l := len(this.stack)
	this.stack = this.stack[:l-1]
	if this.minStack[len(this.minStack)-1] == l-1 {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	// fmt.Println(this.stack, this.minStack)

}

func (this *MinStack) Top() int {
	l := len(this.stack)
	return this.stack[l-1]
}

func (this *MinStack) GetMin() int {
	l := len(this.minStack)
	return this.stack[this.minStack[l-1]]
}
