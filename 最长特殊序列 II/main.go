package main

import "fmt"

func main() {
	fmt.Println(findLUSlength([]string{"aabbcc", "aabbcc", "cb"}))

	//fmt.Println(isSubStr("aabbcc", "aabbcc"))

}

func findLUSlength(strs []string) int {
	max := -1
	for i := 0; i < len(strs); i++ {
		str1 := strs[i]
		isSpecialSub := true
		for j := 0; j < len(strs); j++ {
			if j == i {
				continue
			}
			if isSubStr(str1, strs[j]) {
				isSpecialSub = false
			}
		}
		if isSpecialSub && max < len(str1) {
			max = len(str1)
		}
	}

	return max
}

func isSubStr(str1, str2 string) bool {
	if len(str1) > len(str2) {
		return false
	}

	index1 := 0
	index2 := 0
	for index2 < len(str2) {
		if str1[index1] == str2[index2] {
			index1++
			index2++
			if index1 == len(str1) {
				return true
			}
			continue
		}

		index2++
	}

	return false
}
