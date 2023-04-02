package main

func main() {

}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n, m := len(rowSum), len(colSum)
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, m)
	}

	for i, rSum := range rowSum {
		for j, cSum := range colSum {
			minSum := min(rSum, cSum)
			ans[i][j] = minSum
			rSum -= minSum
			colSum[j] -= minSum
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
