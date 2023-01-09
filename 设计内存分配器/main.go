package main

func main() {

}

type Allocator struct {
	memory []int
}

func Constructor(n int) Allocator {
	return Allocator{
		memory: make([]int, n),
	}
}

func (this *Allocator) Allocate(size int, mID int) int {
	for i := 0; i < len(this.memory); i++ {
		if this.memory[i] != 0 {
			continue
		}
		j := i
		count := 0
		for ; j < len(this.memory) && j < i+size; j++ {
			if this.memory[j] != 0 {
				break
			}
			count++
		}
		if count == size {
			for j := i; j < i+size; j++ {
				this.memory[j] = mID
			}
			return i
		}
		i = j
	}

	return -1
}

func (this *Allocator) Free(mID int) int {
	count := 0
	for i := 0; i < len(this.memory); i++ {
		if this.memory[i] == mID {
			this.memory[i] = 0
			count++
		}
	}
	return count
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.Free(mID);
 */
