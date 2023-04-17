package main

func main() {

}

func rowAndMaximumOnes(mat [][]int) []int {
	mx := 0
	index := 0
	for i := 0; i < len(mat); i++ {
		sum := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				sum++
			}
		}
		if sum > mx {
			mx = sum
			index = i
		}
	}
	return []int{index, mx}
}
