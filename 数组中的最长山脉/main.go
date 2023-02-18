package main

import "fmt"

func main() {
	fmt.Println(longestMountain([]int{1, 2, 3, 4, 5, 4, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5}))
}

func longestMountain(arr []int) int {
	lastUpLen, lastDownLen := -1, -1
	ans := 0
	for i, v := range arr {
		if i == 0 {
			continue
		}

		if v > arr[i-1] {
			if lastDownLen != -1 {
				lastUpLen = -1
				lastDownLen = -1
			}

			if lastUpLen == -1 {
				lastUpLen = 2
			} else {
				lastUpLen++
			}
		}

		if v == arr[i-1] {
			lastUpLen = -1
			lastDownLen = -1
		}

		if v < arr[i-1] {
			if lastDownLen == -1 {
				lastDownLen = 2
			} else {
				lastDownLen++
			}

			if lastUpLen != -1 {
				if ans < lastUpLen+lastDownLen {
					ans = lastUpLen + lastDownLen
				}
			}
		}
	}

	return ans - 1
}
