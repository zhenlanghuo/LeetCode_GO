package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(braceExpansionII("{{a,z},a{b,c},{ab,z}}"))
}

func braceExpansionII(expression string) []string {

	idx := 0

	var expr func() map[string]struct{}
	var term func() map[string]struct{}
	var iterm func() map[string]struct{}

	iterm = func() map[string]struct{} {
		set := make(map[string]struct{})
		if expression[idx] == '{' {
			idx++
			set = expr()
		} else {
			set[expression[idx:idx+1]] = struct{}{}
			//fmt.Println(111)
			//set[fmt.Sprintf("%v", expression[idx])] = struct{}{}
		}
		idx++
		return set
	}

	term = func() map[string]struct{} {
		set := make(map[string]struct{})
		set[""] = struct{}{}
		for idx < len(expression) && ((expression[idx]-'a' >= 0 && expression[idx]-'a' <= 26) || expression[idx] == '{') {
			ans := iterm()
			temp := make(map[string]struct{})
			for k1 := range set {
				for k2 := range ans {
					temp[k1+k2] = struct{}{}
				}
			}
			set = temp
		}
		return set
	}

	expr = func() map[string]struct{} {
		set := make(map[string]struct{})
		for {
			temp := term()
			for key := range temp {
				set[key] = struct{}{}
			}
			if idx < len(expression) && expression[idx] == ',' {
				idx++
				continue
			} else {
				break
			}
		}

		return set
	}

	ansSet := expr()
	//log.Printf("%v", ans)
	fmt.Println(ansSet)

	ans := make([]string, 0, len(ansSet))
	for k := range ansSet {
		ans = append(ans, k)
	}

	sort.Strings(ans)

	return ans
}
