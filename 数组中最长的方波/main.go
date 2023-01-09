package main

import "fmt"

func main() {
	fmt.Println(longestSquareStreak([]int{16 * 16, 3, 6, 16, 4, 8, 2}))
	fmt.Println(longestSquareStreak([]int{81, 2, 3, 5, 6, 7, 9}))
}

func longestSquareStreak(nums []int) int {
	numMap := make(map[int]bool)
	visited := make(map[int]int)
	for _, num := range nums {
		numMap[num] = true
	}

	result := -1
	for _, num := range nums {
		if _, ok := visited[num]; ok {
			continue
		}
		count := 1
		temp := num
		for {
			//fmt.Println(temp)
			if _, ok := numMap[temp*temp]; !ok {
				break
			}
			if _, ok := visited[temp*temp]; ok {
				break
			}
			count++
			visited[temp*temp] = count
			temp = temp * temp
		}
		//fmt.Println(num, temp)
		if _, ok := visited[temp*temp]; ok {
			count = count + visited[temp*temp]
		}
		visited[num] = count
		if count >= 2 && result < count {
			result = count
		}
	}

	//fmt.Println(visited)

	return result
}
