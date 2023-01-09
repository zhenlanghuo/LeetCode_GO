package main

import "fmt"

func main() {
	//fmt.Println(minDeletionSize([]string{"ca", "bb", "ac"}))
	//fmt.Println(minDeletionSize([]string{"xc", "yb", "za"}))
	//fmt.Println(minDeletionSize([]string{"bbjwefkpb", "axmksfchw"}))
	fmt.Println(minDeletionSize([]string{"vdy", "vei", "zvc", "zld"}))
}

func minDeletionSize(strs []string) int {
	n, m := len(strs), len(strs[0])
	bytesList := make([][]byte, n)
	for i := 0; i < n; i++ {
		bytesList[i] = []byte(strs[i])
	}

	isLess := make([][]bool, n)
	for i := 0; i < n; i++ {
		isLess[i] = make([]bool, n)
	}

	count := 0
	prev := -1
	for i := 0; i < m; i++ {
		// 比较第i个字符是否都满足字典序，如果不满足则需要进行删除
		needDelete := false
		for j := 1; j < n; j++ {
			if prev == -1 {
				//fmt.Println(i, j)
				if bytesList[j-1][i] > bytesList[j][i] {
					needDelete = true
					break
				}
			} else {
				if !isLess[j-1][j] && bytesList[j-1][prev] == bytesList[j][prev] && bytesList[j-1][i] > bytesList[j][i] {
					needDelete = true
					break
				}
			}
		}

		if needDelete {
			count++
		} else {
			prev = i
			for j := 1; j < n; j++ {
				if bytesList[j-1][i] < bytesList[j][i] {
					isLess[j-1][j] = true
				}
			}
		}
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
