package main

import "sort"

func main() {

}

func numberOfWeakCharacters(properties [][]int) int {

	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})

	maxDef := 0
	ans := 0
	for _, p := range properties {
		if p[1] < maxDef {
			ans++
		}

		if p[1] > maxDef {
			maxDef = p[1]
		}
	}

	return ans
}
