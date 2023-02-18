package main

func main() {

}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	minI, maxI, i0 := -1, -1, -1
	ans := int64(0)
	for i, num := range nums {
		if num == minK {
			minI = i
		}
		if num == maxK {
			maxI = i
		}
		if num < minK || num > maxK {
			i0 = i
		}
		ans += int64(max(min(minI, maxI)-i0, 0))
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
