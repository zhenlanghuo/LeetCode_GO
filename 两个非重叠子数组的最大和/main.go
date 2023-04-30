package main

func main() {
}

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {

	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	f := func(firstLen int, secondLen int) int {
		maxA := 0
		ans := 0
		for i := firstLen + secondLen - 1; i < n; i++ {
			maxA = max(maxA, sum[i-secondLen+1]-sum[i-secondLen-firstLen+1])
			ans = max(ans, maxA+sum[i+1]-sum[i-secondLen+1])
		}

		return ans
	}

	return max(f(firstLen, secondLen), f(secondLen, firstLen))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
