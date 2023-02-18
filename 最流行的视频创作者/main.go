package main

import "fmt"

func main() {
	fmt.Println(mostPopularCreator([]string{"alice", "bob", "alice", "chris"}, []string{"one", "two", "three", "four"}, []int{5, 10, 5, 4}))
	fmt.Println(mostPopularCreator([]string{"a", "a"}, []string{"aa", "a"}, []int{5, 5}))
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	n := len(creators)
	maxViewId := make(map[string]int)
	maxViewMap := make(map[string]int)
	maxView := 0
	for i := 0; i < n; i++ {
		maxViewMap[creators[i]] += views[i]
		maxView = max(maxView, maxViewMap[creators[i]])
		if _, ok := maxViewId[creators[i]]; !ok {
			maxViewId[creators[i]] = i
		}
		if views[i] > views[maxViewId[creators[i]]] {
			maxViewId[creators[i]] = i
		}

		if views[i] == views[maxViewId[creators[i]]] && ids[i] < ids[maxViewId[creators[i]]] {
			maxViewId[creators[i]] = i
		}
	}

	ans := make([][]string, 0)
	for creator, view := range maxViewMap {
		if view == maxView {
			ans = append(ans, []string{creator, ids[maxViewId[creator]]})
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
