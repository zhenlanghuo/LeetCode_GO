package main

import "fmt"

func main() {
	fmt.Println(cycleLengthQueries(3, [][]int{{5, 3}, {4, 7}, {2, 3}}))
	fmt.Println(cycleLengthQueries(5, [][]int{{8, 17}}))
}

func cycleLengthQueries(n int, queries [][]int) []int {
	result := make([]int, 0, len(queries))
	for _, query := range queries {
		a, b := query[0], query[1]
		count := 0
		for a != b {
			if a > b {
				a = a / 2
			} else {
				b = b / 2
			}
			count++
		}
		result = append(result, count+1)
	}

	return result
}

//func cycleLengthQueries(n int, queries [][]int) []int {
//	result := make([]int, 0)
//	for _, query := range queries {
//		a, b := query[0], query[1]
//		tempA, tempB := a, b
//		countA, countB := 0, 0
//		for tempA != 1 {
//			tempA = tempA / 2
//			countA++
//		}
//		for tempB != 1 {
//			tempB = tempB / 2
//			countB++
//		}
//
//		if countB > countA {
//			a, b = b, a
//			countA, countB = countB, countA
//		}
//
//		fmt.Println(a, b, countA, countB)
//		count := countA - countB
//		for i := 0; i < count; i++ {
//			a = a / 2
//		}
//		if a != b {
//			for i := 0; i < countA; i++ {
//				a = a / 2
//				b = b / 2
//				count += 2
//				if a == b {
//					break
//				}
//			}
//		}
//
//		result = append(result, count+1)
//	}
//
//	return result
//}
