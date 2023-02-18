package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

func partitionLabels(s string) []int {
	byteCount := make(map[byte]int)
	for _, b := range s {
		byteCount[byte(b)]++
	}

	//fmt.Println(byteCount)

	result := make([]int, 0)
	byteCount2 := make(map[byte]bool)
	prev := 0
	for i, b := range s {
		byteCount2[byte(b)] = true
		byteCount[byte(b)]--
		//fmt.Println(byteCount)
		if byteCount[byte(b)] == 0 {
			delete(byteCount2, byte(b))
		}

		if len(byteCount2) == 0 {
			result = append(result, i-prev+1)
			prev = i + 1
		}
	}

	return result
}
