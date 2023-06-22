package main

import "fmt"

func main() {
	fmt.Println(smallestString("cbabc"))
	fmt.Println(smallestString("acbbc"))
	fmt.Println(smallestString("leetcode"))
	fmt.Println(smallestString("aa"))
}

func smallestString(s string) string {

	n := len(s)
	bytes := []byte(s)
	i := 0
	for i < n && bytes[i] == 'a' {
		i++
	}

	if i == n {
		bytes[n-1] = 'z'
	} else {
		for ; i < n && bytes[i] != 'a'; i++ {
			bytes[i] -= 1
		}
	}

	return string(bytes)

}
