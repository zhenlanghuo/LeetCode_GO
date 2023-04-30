package main

func main() {

}

func maximizeSum(nums []int, k int) int {

	mx := nums[0]
	for _, num := range nums {
		mx = max(mx, num)
	}

	ans := 0
	for i := 0; i < k; i++ {
		ans += mx
		mx++
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
