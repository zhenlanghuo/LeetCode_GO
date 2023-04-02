package main

import "fmt"

func main() {
	fmt.Println(findLongestSubarray([]string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7", "H", "I", "J", "K", "L", "M"}))
	fmt.Println(findLongestSubarray([]string{"A", "1"}))
	fmt.Println(findLongestSubarray([]string{"A", "A", "A", "A", "1", "2"}))
}

func findLongestSubarray(array []string) []string {

	n := len(array)
	sums := make([]int, n+1)
	maxLen := 0
	left := 0
	first := make(map[int]int)
	first[0] = -1
	for i, str := range array {
		sums[i+1] = sums[i]
		if []byte(str)[0]-'0' <= 9 {
			sums[i+1]++
		} else {
			sums[i+1]--
		}
		if v, ok := first[sums[i+1]]; ok {
			if maxLen < i-v {
				maxLen = i - v
				left = v
			}
		} else {
			first[sums[i+1]] = i
		}
	}
	//fmt.Println(sums)
	//fmt.Println(left, maxLen)

	return array[left+1 : left+maxLen+1]
}
