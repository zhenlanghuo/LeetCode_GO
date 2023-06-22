package main

func main() {

}

func count(num1 string, num2 string, min_sum int, max_sum int) int {
	const mod int = 1e9 + 7
	f := func(s string) int {
		memory := make([][][]int, len(s))
		for i := 0; i < len(s); i++ {
			memory[i] = make([][]int, max_sum+1)
			for j := 0; j <= max_sum; j++ {
				memory[i][j] = make([]int, 2)
				for k := 0; k < 2; k++ {
					memory[i][j][k] = -1
				}
			}
		}
		var dfs func(p, sum int, limitUp bool) int
		dfs = func(p, sum int, limitUp bool) int {
			if sum > max_sum {
				return 0
			}

			if p == len(s) {
				if sum >= min_sum {
					return 1
				}
				return 0
			}

			switch limitUp {
			case true:
				if memory[p][sum][1] != -1 {
					return memory[p][sum][1]
				}
			case false:
				if memory[p][sum][0] != -1 {
					return memory[p][sum][0]
				}
			}

			up := 9
			if limitUp {
				up = int(s[p] - '0')
			}

			res := 0
			for d := 0; d <= up; d++ {
				res = (res + dfs(p+1, sum+d, limitUp && d == up)) % mod
			}
			switch limitUp {
			case true:
				memory[p][sum][1] = res
			case false:
				memory[p][sum][0] = res
			}

			return res
		}

		return dfs(0, 0, true)
	}

	ans := f(num2) - f(num1) + mod
	sum := 0
	for _, c := range num1 {
		sum += int(c - '0')
	}
	if min_sum <= sum && max_sum >= sum {
		ans++
	}
	return ans % mod
}
