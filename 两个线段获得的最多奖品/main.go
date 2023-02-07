package main

func main() {

}

func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	pre := make([]int, n+1)
	result := 0
	left := 0
	for right, position := range prizePositions {
		for position-prizePositions[left] > k {
			left++
		}
		result = max(result, right-left+1+pre[left])
		pre[right+1] = max(pre[right], right-left+1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
