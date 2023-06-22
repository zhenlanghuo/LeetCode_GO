package main

import "fmt"

func main() {
	fmt.Println(specialPerm([]int{2, 3, 6}))

	//for i := 0; i < 4; i++ {
	//	fmt.Println(1 << uint(i))
	//}
	//fmt.Println(2 | (1 << 2))
}

func specialPerm(nums []int) int {

	mod := int(1e9 + 7)

	n := len(nums)
	size := 1 << uint(n)
	memory := make([][]int, n)
	for i := 0; i < n; i++ {
		memory[i] = make([]int, size)
		for j := 0; j < size; j++ {
			memory[i][j] = -1
		}
	}

	var dfs func(start, mask int) int
	dfs = func(start, mask int) int {

		//fmt.Println(start, mask)

		if mask == size-1 {
			return 1
		}

		if memory[start][mask] != -1 {
			return memory[start][mask]
		}

		ans := 0
		for i := 0; i < n; i++ {
			temp := 1 << uint(i)
			if mask&temp != 0 {
				continue
			}

			if nums[i]%nums[start] == 0 || nums[start]%nums[i] == 0 {
				ans = (ans + dfs(i, mask|temp)) % mod
			}
		}
		memory[start][mask] = ans
		return ans
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = (ans + dfs(i, 1<<(uint(i)))) % mod
	}

	return ans
}
