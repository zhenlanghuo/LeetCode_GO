package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	//fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	//fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))
	//fmt.Println(eraseOverlapIntervals([][]int{{0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}}))
	fmt.Println(eraseOverlapIntervals([][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}))
}

type intervalType struct {
	start int
	end   int
}

type intervalList []intervalType

func (i intervalList) Len() int      { return len(i) }
func (i intervalList) Swap(x, y int) { i[x], i[y] = i[y], i[x] }
func (i intervalList) Less(x, y int) bool {
	if i[x].start == i[y].start {
		return i[x].end < i[y].end
	}
	return i[x].start < i[y].start
}

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	itList := make(intervalList, 0, len(intervals))
	for _, interval := range intervals {
		itList = append(itList, intervalType{start: interval[0], end: interval[1]})
	}

	sort.Sort(itList)

	left := 0
	right := 1
	count := 1
	temp := make(intervalList, 0)
	//temp = append(temp, itList[0])
	for right < n {
		if itList[left].start != itList[right].start {
			if itList[left].end >= itList[right].end {
				//temp = temp[:len(temp)-1]
				//temp = append(temp, itList[right])
				left = right
			} else if itList[left].end <= itList[right].start {
				//temp = append(temp, itList[right])
				left = right
				count++
			}
		}
		right++
	}

	fmt.Println(itList)
	fmt.Println(temp)

	for right < n {

	}

	return n - count
}
