package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canReorderDoubled([]int{-3, -2, -6, -4}))
}

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)
	//fmt.Println(arr)
	end := 0
	for ; end < len(arr); end++ {
		if arr[end] > 0 {
			break
		}
	}
	//fmt.Println(end)
	reverse(arr, 0, end-1)

	arrMap := make(map[int][]int)
	for i, v := range arr {
		if _, ok := arrMap[v]; !ok {
			arrMap[v] = make([]int, 0)
		}
		arrMap[v] = append(arrMap[v], i)
	}

	used := make([]bool, len(arr))

	//fmt.Println(arrMap)
	//fmt.Println(arr)

	for i, v := range arr {
		if used[i] {
			continue
		}
		indexes, ok := arrMap[2*v]
		if !ok || len(indexes) == 0 {
			return false
		}

		arrMap[2*v] = arrMap[2*v][1:]
		used[indexes[0]] = true
		used[i] = true
	}

	return true
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
