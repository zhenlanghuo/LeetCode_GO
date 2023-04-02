package main

func main() {

}

type pair struct {
	x int
	y int
}

func rootCount(edges [][]int, guesses [][]int, k int) int {

	n := len(edges) + 1
	g := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	guessesMap := make(map[pair]bool, len(guesses))
	for _, guess := range guesses {
		guessesMap[pair{x: guess[0], y: guess[1]}] = true
	}

	cnt := 0
	var dfs func(x, fa int)
	dfs = func(x, fa int) {
		for _, child := range g[x] {
			if child == fa {
				continue
			}
			if guessesMap[pair{x: x, y: child}] {
				cnt++
			}
			dfs(child, x)
		}
	}
	dfs(0, -1)

	ans := 0
	var reRoot func(x, fa, cnt int)
	reRoot = func(x, fa, cnt int) {
		if cnt >= k {
			ans++
		}
		for _, child := range g[x] {
			if child == fa {
				continue
			}
			tempCnt := cnt
			if guessesMap[pair{x: x, y: child}] {
				tempCnt--
			}
			if guessesMap[pair{x: child, y: x}] {
				tempCnt++
			}
			reRoot(child, x, tempCnt)
		}
	}
	reRoot(0, -1, cnt)

	return ans
}
