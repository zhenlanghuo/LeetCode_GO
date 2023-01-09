package main

import "fmt"

func main() {
	//rand.Seed(time.Now().Unix())
	//arr := make([]*pair, 10)
	//for i := 0; i < len(arr); i++ {
	//	arr[i] = &pair{val: rand.Intn(100), index: i}
	//}
	//
	//temp = make([]*pair, len(arr))
	//sort(arr, 0, len(arr)-1)
	//fmt.Println(arr)

	fmt.Println(countSmaller([]int{1, -1}))
}

var temp []*pair
var count map[int]int

type pair struct {
	val   int
	index int
}

func countSmaller(nums []int) []int {
	temp = make([]*pair, len(nums))
	count = make(map[int]int)
	pairs := make([]*pair, len(nums))
	for i := 0; i < len(nums); i++ {
		pairs[i] = &pair{val: nums[i], index: i}
	}

	sort(pairs, 0, len(pairs)-1)

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = count[i]
	}

	return result
}

func sort(nums []*pair, lo, hi int) {
	if lo == hi {
		return
	}
	mid := (lo + hi) / 2
	sort(nums, lo, mid)
	sort(nums, mid+1, hi)
	merge(nums, lo, hi, mid)
}

func merge(nums []*pair, lo, hi, mid int) {
	for i := lo; i <= hi; i++ {
		temp[i] = nums[i]
	}

	l, r := lo, mid+1
	for lo <= hi {
		if l > mid {
			nums[lo] = temp[r]
			r++
		} else if r > hi {
			nums[lo] = temp[l]
			count[temp[l].index] += r - mid - 1
			l++
		} else {
			if temp[l].val > temp[r].val {
				nums[lo] = temp[r]
				r++
			} else {
				nums[lo] = temp[l]
				count[temp[l].index] += r - mid - 1
				l++
			}
		}
		lo++
	}
}
