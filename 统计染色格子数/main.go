package main

import "fmt"

func main() {
	fmt.Println(coloredCells(3))
}

func coloredCells(n int) int64 {
	ans := int64(1)
	for i := 1; i < n; i++ {
		//fmt.Println(ans)
		ans += int64(i * 4)
	}
	return ans
}
