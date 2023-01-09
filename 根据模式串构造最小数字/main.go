package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallestNumber("IIIDIDDD"))
	//fmt.Println(smallestNumber("DDD"))
	//fmt.Println(smallestNumber("I"))
}

func smallestNumber(pattern string) string {
	// 贪心
	// 从左到右找出连续的先I后D的子串
	// 比如 IIIDD, 然后拿出目前最小的6个数字(n+1), 比如 123456, 对后面三个数字（DD的数量加+1）就行reverse(123654)
	n := len(pattern)
	bytes := make([]byte, len(pattern)+1)
	for i := 0; i < len(bytes); i++ {
		bytes[i] = byte(i+1) + '0'
	}

	i := 0
	for i < n {
		j := i
		ICount := 0
		for ; j < n && pattern[j] == 'I'; j++ {
			ICount++
		}
		for ; j < n && pattern[j] == 'D'; j++ {
		}

		reverse(bytes, i+ICount, j)
		i = j
	}

	return string(bytes)
}

func reverse(bytes []byte, start, end int) {
	for start < end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start++
		end--
	}
}

//func smallestNumber(pattern string) string {
//
//	used := make([]bool, 9)
//
//	var dfs func(temp string)
//
//	result := ""
//	for i := 9; i >= 1; i-- {
//		if len(result) == len(pattern)+1 {
//			break
//		}
//		result += strconv.FormatUint(uint64(i), 10)
//	}
//	//fmt.Println(result)
//
//	dfs = func(temp string) {
//		if len(temp) == len(pattern)+1 {
//			if result > temp {
//				result = temp
//			}
//			return
//		}
//
//		for i := 1; i <= 9; i++ {
//			if used[i-1] {
//				continue
//			}
//			if len(temp) != 0 {
//				lastNum, _ := strconv.ParseUint(temp[len(temp)-1:], 10, 64)
//				switch pattern[len(temp)-1] {
//				case 'I':
//					if lastNum > uint64(i) {
//						continue
//					}
//				case 'D':
//					if lastNum < uint64(i) {
//						continue
//					}
//				}
//
//			}
//
//			used[i-1] = true
//			dfs(temp + strconv.FormatUint(uint64(i), 10))
//			used[i-1] = false
//		}
//	}
//
//	dfs("")
//
//	return result
//}
