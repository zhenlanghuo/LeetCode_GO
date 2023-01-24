package main

func main() {

}

func rangeAddQueries(n int, queries [][]int) [][]int {
	diffs := make([][]int, n)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		diffs[i] = make([]int, n)
		mat[i] = make([]int, n)
	}

	for _, query := range queries {
		start, end := query[1], query[3]
		for i := query[0]; i <= query[2]; i++ {
			diffAdd(diffs[i], 1, start, end)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				mat[i][j] = diffs[i][j]
			} else {
				mat[i][j] = mat[i][j-1] + diffs[i][j]
			}
		}
	}

	return mat
}

func diffAdd(diff []int, val, start, end int) {
	diff[start] += val
	if end+1 < len(diff) {
		diff[end+1] -= val
	}
}
