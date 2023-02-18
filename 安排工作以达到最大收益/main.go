package main

import "sort"

func main() {

}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {

	difficultyIndex := make([]int, len(difficulty))
	for i := 0; i < len(difficulty); i++ {
		difficultyIndex[i] = i
	}
	sort.Slice(difficultyIndex, func(i, j int) bool {
		return difficulty[difficultyIndex[i]] < difficulty[difficultyIndex[j]]
	})

	sort.Ints(worker)

	tempMax, ans, i := 0, 0, 0
	for _, w := range worker {
		for i < len(difficultyIndex) && difficulty[difficultyIndex[i]] <= w {
			tempMax = max(tempMax, profit[difficultyIndex[i]])
			i++
		}

		ans += tempMax
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
