package main

import "strings"

func main() {

}

func isValidSerialization(preorder string) bool {

	splits := strings.Split(preorder, ",")
	n := len(splits)

	nullCount := 0

	for i := n - 1; i >= 0; i-- {
		if splits[i] == "#" {
			nullCount++
		} else {
			if nullCount < 2 {
				return false
			}
			nullCount--
		}
	}

	if nullCount != 0 {
		return false
	}

	return true
}
