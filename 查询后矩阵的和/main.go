package main

func main() {

}

func matrixSumQueries(n int, queries [][]int) int64 {
	ans := int64(0)
	row := make(map[int]bool)
	col := make(map[int]bool)
	for i := len(queries) - 1; i >= 0; i-- {
		t, index, val := queries[i][0], queries[i][1], queries[i][2]
		switch t {
		case 0:
			if !row[index] {
				ans += int64(val) * int64(n-len(col))
			}
			row[index] = true
		case 1:
			if !col[index] {
				ans += int64(val) * int64(n-len(row))
			}
			col[index] = true
		}
	}

	return ans
}
