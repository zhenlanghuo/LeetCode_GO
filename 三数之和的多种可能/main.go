package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum([]int{1, 1, 1, 3, 2, 2}, 3))
}

func threeSumMulti(arr []int, target int) int {
	sort.Ints(arr)
	ans := 0
	mod := int(1e9 + 7)
	for i, v := range arr {
		ans = (ans + twoSum(arr[i+1:], target-v)) % mod
	}
	return ans
}

func twoSum(arr []int, target int) int {
	n := len(arr)
	ans := 0
	l, r := 0, n-1
	for l < r {
		//fmt.Println(l, r, -1)
		if arr[l]+arr[r] == target {
			countl, countr := 0, 0
			numl, numr := arr[l], arr[r]
			for l < n && arr[l] == numl {
				countl++
				l++
			}

			for r >= 0 && arr[r] == numr {
				countr++
				r--
			}
			//fmt.Println(l, r)
			if numl == numr {
				ans += countl * (countl - 1) / 2
			} else {
				ans += countl * countr
			}
		} else if arr[l]+arr[r] > target {
			r--
		} else {
			l++
		}
	}

	return ans
}
