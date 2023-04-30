package main

import "fmt"

func main() {
	fmt.Println(challengeOfTheKeeper([]string{"........T.", "..........", "....#..#..", "........#.", "..#......#", "#...#.....", ".........#", ".#S.......", "#.........", "..##.#..#."}))
}

func challengeOfTheKeeper(maze []string) int {
	n, m := len(maze), len(maze[0])
	disFromT := make([][]int, n)
	for i := 0; i < n; i++ {
		disFromT[i] = make([]int, m)
		for j := 0; j < m; j++ {
			disFromT[i][j] = -2
		}
	}

	type Point struct {
		x int
		y int
	}
	tPoint := &Point{}
	sPoint := &Point{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if maze[i][j] == 'T' {
				tPoint.x = i
				tPoint.y = j
			}

			if maze[i][j] == 'S' {
				sPoint.x = i
				sPoint.y = j
			}
		}
	}

	directs := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	q := []*Point{tPoint}
	disFromT[tPoint.x][tPoint.y] = 0
	for len(q) != 0 {
		temp := q
		q = nil
		for _, p := range temp {
			x, y := p.x, p.y
			for _, direct := range directs {
				nextX, nextY := x+direct[0], y+direct[1]
				if nextX >= 0 && nextX < n && nextY >= 0 && nextY < m && disFromT[nextX][nextY] == -2 {
					if maze[nextX][nextY] == '#' {
						disFromT[nextX][nextY] = -1
					} else {
						disFromT[nextX][nextY] = disFromT[x][y] + 1
						q = append(q, &Point{x: nextX, y: nextY})
					}
				}
			}
		}
	}

	fmt.Println(disFromT)

	if disFromT[sPoint.x][sPoint.y] == -2 {
		return -1
	}

	vis := make([][]bool, n)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(i, j, maxDis int) bool
	dfs = func(i, j, maxDis int) bool {
		if i < 0 || i >= n || j < 0 || j >= m || vis[i][j] || maze[i][j] == '#' {
			return false
		}

		if maze[i][j] == 'T' {
			return true
		}
		vis[i][j] = true
		if maze[i][j] == '.' {
			points := []*Point{&Point{x: i, y: m - 1 - j}, {x: n - 1 - i, y: j}}
			for _, p := range points {
				if maze[p.x][p.y] != '#' && (disFromT[p.x][p.y] > maxDis || disFromT[p.x][p.y] == -2) {
					return false
				}
			}
		}
		for _, dir := range dirs {
			if dfs(i+dir[0], j+dir[1], maxDis) {
				return true
			}
		}
		return false
	}

	check := func(maxDis int) bool {
		for i := 0; i < n; i++ {
			vis[i] = make([]bool, m)
		}
		return dfs(sPoint.x, sPoint.y, maxDis)
	}

	left, right := 0, n*m
	for left <= right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left >= m*n {
		return -1
	}

	return left
}
