package main

import "math"

func main() {

}

func minimumCost(start []int, target []int, specialRoads [][]int) int {
	type pair struct {
		x int
		y int
	}
	t := pair{target[0], target[1]}
	dis := make(map[pair]int)
	vis := make(map[pair]bool)
	dis[t] = math.MaxInt64
	dis[pair{start[0], start[1]}] = 0
	for {
		v, dv := pair{}, -1
		for p, d := range dis {
			if !vis[p] && (dv < 0 || d < dv) {
				v = p
				dv = d
			}
		}
		if v == t {
			return dv
		}
		vis[v] = true
		dis[t] = min(dis[t], dv+abs(t.x-v.x)+abs(t.y-v.y))
		for _, r := range specialRoads {
			a := pair{r[0], r[1]}
			b := pair{r[2], r[3]}
			cost := r[4]
			d := min(dv+abs(b.x-v.x)+abs(b.y-v.y), dv+abs(a.x-v.x)+abs(a.y-v.y)+cost)
			//d := dv + abs(a.x-v.x) + abs(a.y-v.y) + cost
			if _, ok := dis[b]; !ok {
				dis[b] = d
			}
			dis[b] = min(dis[b], d)
		}
		for _, r := range specialRoads {
			w := pair{r[2], r[3]}
			d := dv + abs(r[0]-v.x) + abs(r[1]-v.y) + r[4]
			if dw, ok := dis[w]; !ok || d < dw {
				dis[w] = d
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
