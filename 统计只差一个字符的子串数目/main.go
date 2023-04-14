package main

import "fmt"

func main() {
	//fmt.Println(countSubstrings("aba", "baba"))
	fmt.Println(countSubstrings("abe", "bbc"))
	fmt.Println(countSubstrings("ab", "bb"))
	fmt.Println(countSubstrings("abbab", "bbbbb"))
}

func countSubstrings(s string, t string) int {
	m := make(map[byte][]int)
	for i, b := range t {
		m[byte(b)] = append(m[byte(b)], i)
	}

	fmt.Println(m)

	ans := 0
	for i, b := range s {
		ans += len(t) - len(m[byte(b)])
		for _, start := range m[byte(b)] {
			step := 0
			for i+step < len(s) && start+step < len(t) && s[i+step] == t[start+step] {
				step++
			}
			if i+step < len(s) && start+step < len(t) && s[i+step] != t[start+step] {
				//fmt.Println(b, i, start, step)
				ans++
			}
			step++
			for i+step < len(s) && start+step < len(t) && s[i+step] == t[start+step] {
				step++
				ans++
			}

			step = 0
			for i-step >= 0 && start-step >= 0 && s[i-step] == t[start-step] {
				step++
			}
			if i-step >= 0 && start-step >= 0 && s[i-step] != t[start-step] {
				//fmt.Println(b, i, start, step)
				ans++
			}
		}
	}

	return ans
}
