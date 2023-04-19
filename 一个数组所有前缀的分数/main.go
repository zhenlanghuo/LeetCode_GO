package main

func main() {

}

func findPrefixScore(nums []int) []int64 {
	n := len(nums)
	ans := make([]int64, n)
	ans[0] = int64(nums[0]) * 2
	mx := nums[0]
	for i := 1; i < n; i++ {
		mx = max(mx, nums[i])
		ans[i] = ans[i-1] + int64(nums[i]+mx)

	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
