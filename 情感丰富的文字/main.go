package main

import "fmt"

func main() {
	//fmt.Println(expressiveWords("heeelloooe", []string{"hello", "hi", "helo"}))
	fmt.Println(expressiveWords("dddiiiinnssssssoooo", []string{"dinnssoo", "ddinso", "ddiinnso", "ddiinnssoo", "ddiinso", "dinsoo", "ddiinsso", "dinssoo", "dinso"}))
	//fmt.Println(expressiveWords("dddiiiinnssssssoooo", []string{"ddiinnso"}))
}

func expressiveWords(s string, words []string) int {
	ans := 0
	for _, word := range words {
		i, j := 0, 0
		flag := true
		for i < len(s) && j < len(word) {
			if s[i] != word[j] {
				flag = false
				break
			}

			ch := s[i]
			cnti, cntj := 0, 0
			for i < len(s) && s[i] == ch {
				i++
				cnti++
			}

			for j < len(word) && word[j] == ch {
				j++
				cntj++
			}

			if cnti < cntj || (cnti > cntj && cnti < 3) {
				flag = false
				break
			}
		}
		//fmt.Println(word, flag, j == len(word))
		if flag && j == len(word) && i == len(s) {
			ans++
		}
	}

	return ans
}
