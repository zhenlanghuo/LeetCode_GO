package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
}

type interval struct {
	interval []int
}

type intervalSlice []interval

func (s intervalSlice) Len() int {
	return len(s)
}

func (s intervalSlice) Less(i, j int) bool {
	return s[i].interval[0] < s[j].interval[0]
}

func (s intervalSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func merge(intervals [][]int) [][]int {

	intervalList := make([]interval, 0, len(intervals))
	for _, interval_ := range intervals {
		intervalList = append(intervalList, interval{interval: interval_})
	}

	sort.Sort(intervalSlice(intervalList))

	result := make([][]int, 0)
	for i := 0; i < len(intervalList); i++ {
		temp := intervalList[i].interval
		for i+1 < len(intervalList) {
			if intervalList[i+1].interval[0] > temp[1] {
				break
			}

			temp[1] = max(intervalList[i+1].interval[1], temp[1])
			i++
		}
		result = append(result, temp)
	}

	return result
}
