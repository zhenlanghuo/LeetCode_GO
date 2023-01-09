package main

func main() {

}

type DataStream struct {
	value int
	k     int
	now   int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		value: value,
		k:     k,
	}
}

func (this *DataStream) Consec(num int) bool {
	if num == this.value {
		if this.now < this.k {
			this.now++
		}
	} else {
		this.now = 0
	}

	return this.now == this.k
}
