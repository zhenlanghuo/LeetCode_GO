package 两点之间不包含任何点的最宽垂直区域

import "sort"

func main() {

}

func maxWidthOfVerticalArea(points [][]int) int {

	ans := 0
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		ans = max(ans, points[i][0]-points[i-1][0])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
