package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"))
	fmt.Println(lengthLongestPath("dir\n        file.txt"))
	//fmt.Println(lengthLongestPath("dir\\n\\tsubdir1\\n\\t\\tfile1.ext\\n\\t\\tsubsubdir1\\n\\tsubdir2\\n\\t\\tsubsubdir2\\n\\t\\t\\tfile2.ext"))
	//fmt.Println(lengthLongestPath("a"))
	//fmt.Println(lengthLongestPath("file1.txt\\nfile2.txt\\nlongfile.txt"))
}

var i = 0

func lengthLongestPath(input string) int {
	i = 0
	getFileOrDirName := func(input string) string {
		tempI := i
		for tempI < len(input) {
			switch input[tempI] {
			case '\n':
				return input[i:tempI]
			default:
				tempI++
			}
		}
		return input[i:]
	}

	ans := 0
	stack := make([]string, 0)
	level := 0
	curLen := 0
	for i < len(input) {
		switch input[i] {
		case '\n':
			i++
			level = 0
			for input[i] == '\t' {
				i++
				level++
			}
		default:
			fileOrDir := getFileOrDirName(input)
			i += len(fileOrDir)
			if strings.Contains(fileOrDir, ".") {
				curLen = 0
				for j := 0; j < level; j++ {
					curLen += len(stack[j]) + 1
				}
				//fmt.Println(stack, level, fileOrDir, curLen)
				ans = max(ans, curLen+len(fileOrDir))
			} else {
				//fmt.Println(stack, fileOrDir, curLen, level)
				stack = stack[:level]
				stack = append(stack, fileOrDir)
				//curLen = 0
				//for _, v := range stack {
				//	curLen += len(v) + 1
				//}
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
