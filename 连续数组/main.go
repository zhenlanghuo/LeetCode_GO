package main

func main() {

}

func findMaxLength(nums []int) int {
	n := len(nums)
	sums := make([]int, n+1)
	maxLen := 0
	first := make(map[int]int)
	first[0] = -1
	for i, num := range nums {
		sums[i+1] = sums[i]
		if num == 1 {
			sums[i+1]++
		} else {
			sums[i+1]--
		}
		if v, ok := first[sums[i+1]]; ok {
			if maxLen < i-v {
				maxLen = i - v
			}
		} else {
			first[sums[i+1]] = i
		}
	}
	return maxLen
}
