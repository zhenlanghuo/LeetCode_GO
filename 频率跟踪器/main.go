package main

func main() {

}

type FrequencyTracker struct {
	m1 map[int]int
	m2 map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		m1: make(map[int]int),
		m2: make(map[int]int),
	}
}

func (this *FrequencyTracker) Add(number int) {
	if this.m1[number] != 0 {
		this.m2[this.m1[number]]--
	}
	this.m1[number]++
	this.m2[this.m1[number]]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.m1[number] == 0 {
		return
	}
	this.m2[this.m1[number]]--
	this.m1[number]--
	this.m2[this.m1[number]]++
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.m2[frequency] > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
