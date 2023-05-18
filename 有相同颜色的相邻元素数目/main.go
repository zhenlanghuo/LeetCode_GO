package main

func main() {
	//fmt.Println(colorTheArray(4))
}

func colorTheArray(n int, queries [][]int) []int {
	ans := make([]int, len(queries))

	colors := make([]int, n)

	for i := 0; i < len(queries); i++ {
		index, color := queries[i][0], queries[i][1]
		preColor := colors[index]
		colors[index] = color

		if i == 0 {
			continue
		}

		ans[i] = ans[i-1]
		if preColor != 0 {
			if index-1 >= 0 && colors[index-1] == preColor {
				ans[i]--
			}

			if index+1 < n && colors[index+1] == preColor {
				ans[i]--
			}
		}

		if index-1 >= 0 && colors[index-1] == color {
			ans[i]++
		}

		if index+1 < n && colors[index+1] == color {
			ans[i]++
		}
	}

	return ans
}
