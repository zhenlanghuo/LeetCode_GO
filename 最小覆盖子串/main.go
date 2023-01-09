package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

//func minWindow(s string, t string) string {
//	byteCountMap := make(map[byte]int)
//	for i := 0; i < len(t); i++ {
//		if _, ok := byteCountMap[t[i]]; !ok {
//			byteCountMap[t[i]] = 0
//		}
//		byteCountMap[t[i]]++
//	}
//
//	indexRecord := make([]bool, len(s))
//	minIndex := -1
//	maxIndex := 0
//	count := 0
//
//	result := ""
//
//	byteIndexMap := make(map[byte][]int)
//	for i := 0; i < len(s); i++ {
//		_, ok := byteCountMap[s[i]]
//		if !ok {
//			continue
//		}
//
//		_, ok = byteIndexMap[s[i]]
//		if !ok {
//			byteIndexMap[s[i]] = []int{}
//		}
//		byteIndexMap[s[i]] = append(byteIndexMap[s[i]], i)
//
//		if minIndex == -1 {
//			minIndex = i
//		}
//
//		maxIndex = i
//		indexRecord[i] = true
//
//		if len(byteIndexMap[s[i]]) > byteCountMap[s[i]] {
//			tempIndex := byteIndexMap[s[i]][0]
//			if tempIndex == minIndex {
//				nextMinIndex := minIndex + 1
//				for indexRecord[nextMinIndex] == false {
//					nextMinIndex++
//				}
//				minIndex = nextMinIndex
//			}
//			indexRecord[tempIndex] = false
//			byteIndexMap[s[i]] = byteIndexMap[s[i]][1:]
//		} else {
//			count++
//		}
//
//		if count == len(t) && (result == "" || len(result) > maxIndex-minIndex+1) {
//			result = s[minIndex : maxIndex+1]
//		}
//
//	}
//
//	return result
//}

func minWindow(s string, t string) string {

	//byteMap := make(map[byte]bool)
	cnt := make([]int, 256)
	for _, b := range t {
		//byteMap[byte(b)] = true
		cnt[b]--
	}
	matchTotal := 0

	result := ""
	left := 0
	for right, b := range s {
		cnt[b]++
		if cnt[b] <= 0 {
			matchTotal++
		}

		for matchTotal == len(t) && left <= right {
			if matchTotal == len(t) && (len(result) > right-left+1 || len(result) == 0) {
				result = s[left : right+1]
			}
			cnt[s[left]]--
			if cnt[s[left]] < 0 {
				matchTotal--
			}
			left++
		}
	}

	return result
}
