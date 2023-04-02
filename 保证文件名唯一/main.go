package main

import "fmt"

func main() {
	fmt.Println(getFolderNames([]string{"gta", "gta(1)", "gta", "avalon"}))
	fmt.Println(getFolderNames([]string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"}))
	fmt.Println(getFolderNames([]string{"kaido", "kaido(1)", "kaido", "kaido(1)"}))
}

func getFolderNames(names []string) []string {
	m := make(map[string]int)
	ans := make([]string, 0, len(names))
	for _, name := range names {
		if m[name] == 0 {
			m[name] = 1
			ans = append(ans, name)
		} else {
			for m[fmt.Sprintf("%s(%d)", name, m[name])] != 0 {
				m[name]++
			}
			newName := fmt.Sprintf("%s(%d)", name, m[name])
			ans = append(ans, newName)
			m[newName] = 1
		}
	}

	return ans
}
