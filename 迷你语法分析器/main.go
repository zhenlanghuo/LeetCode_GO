package main

import (
	"fmt"
	"strconv"
)

func main() {

	bytes := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	fmt.Println(bytes)
}

type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool {
	return true
}

func (this NestedInteger) GetInteger() int {
	return 0
}

func (n *NestedInteger) SetInteger(value int) {}

func (this *NestedInteger) Add(elem NestedInteger) {}

func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

var i = 0

func deserialize(s string) *NestedInteger {
	i = 0
	getNumString := func(s string) string {
		tempI := i
		for tempI < len(s) {
			if s[tempI]-'0' <= 9 || s[tempI] == '-' {
				tempI++
			} else {
				return s[i:tempI]
			}
		}
		return s[i:]
	}

	var dfs func(s string) *NestedInteger
	dfs = func(s string) *NestedInteger {
		ni := &NestedInteger{}
		if s[i] == '[' {
			i++
		} else {
			numStr := getNumString(s)
			i += len(numStr)
			num, _ := strconv.ParseInt(numStr, 10, 64)
			//fmt.Println(numStr, num)
			ni.SetInteger(int(num))
			return ni
		}

		for i < len(s) {
			switch s[i] {
			case '[':
				ni.Add(*dfs(s))
			case ']':
				i++
				return ni
			case ',':
				i++
			default:
				ni.Add(*dfs(s))
			}
		}

		return ni
	}

	return dfs(s)
}
