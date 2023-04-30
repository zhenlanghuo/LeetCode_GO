package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(adventureCamp([]string{"leet->code", "leet->code->Campsite->Leet", "leet->code->leet->courier"}))
	fmt.Println(adventureCamp([]string{"Alice->Dex", "", "Dex"}))
	fmt.Println(adventureCamp([]string{"", "Gryffindor->Slytherin->Gryffindor", "Hogwarts->Hufflepuff->Ravenclaw"}))
}

func adventureCamp(expeditions []string) int {
	ans := -1
	mx := 0
	m := make(map[string]bool)
	for i, expedition := range expeditions {
		count := 0
		splits := strings.Split(expedition, "->")
		for _, split := range splits {
			if strings.TrimSpace(split) != "" && m[split] == false {
				count++
				m[split] = true
			}
		}
		//fmt.Println(i, count, m, splits)
		if i != 0 {
			if count > mx {
				ans = i
				mx = count
			}
		}
	}

	return ans
}
