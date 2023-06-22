package main

func main() {

}

type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool {
	return true
}

func (this NestedInteger) GetInteger() int {
	return 0
}

func (n *NestedInteger) SetInteger(value int) {}

func (this *NestedInteger) Add(elem NestedInteger) {}

func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

type NestedIterator struct {
	stack [][]*NestedInteger
	next  int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		stack: [][]*NestedInteger{nestedList},
	}
}

func (this *NestedIterator) Next() int {
	return this.next
}

func (this *NestedIterator) HasNext() bool {
	if len(this.stack) != 0 {
		if len(this.stack[len(this.stack)-1]) != 0 {
			if this.stack[len(this.stack)-1][0].IsInteger() {
				this.next = this.stack[len(this.stack)-1][0].GetInteger()
				this.stack[len(this.stack)-1] = this.stack[len(this.stack)-1][1:]
			} else {
				tempList := this.stack[len(this.stack)-1][0].GetList()
				this.stack[len(this.stack)-1] = this.stack[len(this.stack)-1][1:]
				this.stack = append(this.stack, tempList)
				return this.HasNext()
			}
		} else {
			this.stack = this.stack[:len(this.stack)-1]
			return this.HasNext()
		}
	}

	return false
}
