package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println(minimumTotalDistance([]int{0, 4, 6}, [][]int{{2, 2}, {6, 2}}))
	fmt.Println(minimumTotalDistance([]int{9, 11, 99, 101}, [][]int{{10, 1}, {7, 1}, {14, 1}, {100, 1}, {96, 1}, {103, 1}}))
	fmt.Println(math.MaxInt64)
	//[9, 11, 99, 101]
	//	[[10, 1], [7, 1], [14, 1], [100, 1], [96, 1],[103, 1]]
}

func minimumTotalDistance(robot []int, factory [][]int) int64 {

	sort.Slice(robot, func(i, j int) bool {
		return robot[i] < robot[j]
	})

	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	n, m := len(robot), len(factory)
	memory := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		memory[i] = make([]int64, m+1)
		for j := 0; j <= m; j++ {
			memory[i][j] = -1
		}
	}

	var dfs func(i, j int) int64
	dfs = func(i, j int) int64 {
		ans := int64(math.MaxInt64)
		//ans := int64(1e18)
		if i >= n {
			// 没有机器需要修理了
			return 0
		}

		if memory[i][j] != -1 {
			return memory[i][j]
		}

		if j == m-1 {
			// 最后一个工厂, 如果工厂修理的机器人数目少于剩余的机器人数目, 返回最大值
			if factory[j][1] < n-i {
				return math.MaxInt64
				//return 1e18
			}
			sum := int64(0)
			for _, x := range robot[i:] {
				sum += abs(int64(x - factory[j][0]))
			}
			memory[i][j] = sum
			return sum
		}

		if j == m {
			return math.MaxInt64
			//return 1e18
		}

		temp := dfs(i, j+1) // 一个都不修
		if temp != math.MaxInt64 {
			ans = min(ans, temp)
		}
		sum := int64(0)
		for k := 0; k < factory[j][1]; k++ {
			if i+k < n {
				sum += abs(int64(robot[i+k] - factory[j][0]))
				temp = dfs(i+k+1, j+1)
				if temp == math.MaxInt64 {
					continue
				}
				ans = min(ans, temp+sum)
			}
		}

		memory[i][j] = ans
		return ans
	}

	return dfs(0, 0)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
