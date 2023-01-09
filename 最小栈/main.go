package main

func main() {

}

//type MinStack struct {
//	stack   []int
//	minList *list.List
//}
//
//func Constructor() MinStack {
//	return MinStack{
//		stack:   make([]int, 0),
//		minList: list.New(),
//	}
//}
//
//func (this *MinStack) Push(val int) {
//	this.stack = append(this.stack, val)
//	if this.minList.Len() == 0 || this.stack[this.minList.Back().Value.(int)] > val {
//		this.minList.PushBack(len(this.stack) - 1)
//	}
//}
//
//func (this *MinStack) Pop() {
//	if this.minList.Back().Value.(int) == len(this.stack)-1 {
//		this.minList.Remove(this.minList.Back())
//	}
//	this.stack = this.stack[:len(this.stack)-1]
//}
//
//func (this *MinStack) Top() int {
//	return this.stack[len(this.stack)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	return this.stack[this.minList.Back().Value.(int)]
//}

type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   0,
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.min = val
		return
	}

	this.stack = append(this.stack, val-this.min)
	if val < this.min {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	top := this.stack[len(this.stack)-1]
	if top <= 0 {
		this.min = this.min - top
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	top := this.stack[len(this.stack)-1]
	if top <= 0 {
		return this.min
	}
	return top + this.min
}

func (this *MinStack) GetMin() int {
	return this.min
}
