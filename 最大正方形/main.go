package main

func main() {

}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, 0, len(matrix)+1)
	for i := 0; i < len(matrix)+1; i++ {
		dp = append(dp, make([]int, len(matrix[0])+1))
	}

	byte1 := []byte("1")[0]
	max := 0
	for i := 1; i < len(matrix)+1; i++ {
		for j := 1; j < len(matrix[0])+1; j++ {
			if matrix[i-1][j-1] != byte1 {
				continue
			}

			min_ := min(dp[i-1][j], dp[i][j-1])
			min_ = min(min_, dp[i-1][j-1])
			dp[i][j] = min_ + 1
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	return max * max
}