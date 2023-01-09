package main

import "fmt"

func main() {
	fmt.Println(isRectangleCover([][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}}))
}

func isRectangleCover(rectangles [][]int) bool {
	sum := 0
	minX, minY, maxX, maxY := rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	type point struct{ x, y int }
	pointMap := make(map[point]int)
	for _, rectangle := range rectangles {
		x1, y1, x2, y2 := rectangle[0], rectangle[1], rectangle[2], rectangle[3]
		sum += (x2 - x1) * (y2 - y1)

		minX = min(minX, x1)
		minY = min(minY, y1)
		maxX = max(maxX, x2)
		maxY = max(maxY, y2)

		pointMap[point{x1, y1}]++
		pointMap[point{x1, y2}]++
		pointMap[point{x2, y1}]++
		pointMap[point{x2, y2}]++
	}

	if sum != (maxX-minX)*(maxY-minY) {
		return false
	}

	if pointMap[point{minX, minY}] != 1 || pointMap[point{minX, maxY}] != 1 ||
		pointMap[point{maxX, minY}] != 1 || pointMap[point{maxX, maxY}] != 1 {
		return false
	}

	delete(pointMap, point{minX, minY})
	delete(pointMap, point{minX, maxY})
	delete(pointMap, point{maxX, minY})
	delete(pointMap, point{maxX, maxY})

	for _, v := range pointMap {
		if v != 2 && v != 4 {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
