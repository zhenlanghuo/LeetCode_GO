package main

func main() {

}

func maximumCount(nums []int) int {
	neg, pos := 0, 0

	for _, num := range nums {
		if num > 0 {
			pos++
		}
		if num < 0 {
			neg++
		}
	}

	if neg > pos {
		return neg
	}
	return pos
}
