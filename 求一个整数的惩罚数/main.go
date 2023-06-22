package main

import "fmt"

func main() {
	fmt.Println(punishmentNumber(37))
}

func punishmentNumber(n int) int {
	var dfs func(ii, target, cur int) bool
	dfs = func(ii, target, cur int) bool {
		if cur > target {
			return false
		}

		if cur+ii == target {
			return true
		}

		for k := 10; ii/k > 0; k *= 10 {
			if dfs(ii%k, target, cur+ii/k) {
				return true
			}
		}

		return false
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if dfs(i*i, i, 0) {
			//fmt.Println(i)
			ans += i * i
		}
	}

	return ans
}
