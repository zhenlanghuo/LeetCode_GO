package main

import "fmt"

func main() {
	array := []*Score{
		&Score{name: "a", yuwen: 96, shuxue: 95},
		&Score{name: "b", yuwen: 97, shuxue: 95},
		&Score{name: "c", yuwen: 95, shuxue: 97},
		&Score{name: "d", yuwen: 94, shuxue: 95},
		&Score{name: "e", yuwen: 94, shuxue: 95},
		&Score{name: "f", yuwen: 94, shuxue: 95},
	}
	sort(array, 0, len(array)-1)
	for _, v := range array {
		fmt.Println(v)
	}
}

type Score struct {
	name   string
	yuwen  int
	shuxue int
}

func (s *Score) LessThan(a *Score) bool {

	if s.yuwen+s.shuxue > a.yuwen+a.shuxue {
		return true
	}

	if s.yuwen+s.shuxue < a.yuwen+a.shuxue {
		return false
	}

	if s.shuxue > a.shuxue {
		return true
	}

	if s.shuxue < a.shuxue {
		return false
	}

	if s.name > a.name {
		return true
	}

	return false
}

func sort(array []*Score, start, end int) {
	if end-start <= 1 {
		return
	}
	part := partion(array, start, end)
	fmt.Println(part)
	print(array)
	sort(array, start, part-1)
	sort(array, part+1, end)
}

func print(array []*Score) {
	str := ""
	for _, v := range array {
		str += fmt.Sprintf("%s", *v)
	}
	fmt.Println(str)
}

//func partion(array []*Score, i, j int) int {
//	piovt := array[i]
//	left := i
//	right := j
//	for left != right {
//		for left < right && piovt.LessThan(array[right]) {
//			right--
//		}
//
//		for left < right && !piovt.LessThan(array[left]) {
//			left++
//		}
//
//		if left < right {
//			array[left], array[right] = array[right], array[left]
//		}
//	}
//
//	print(array)
//	//fmt.Println(i, left)
//	array[i], array[left] = array[left], array[i]
//	return left
//}

func partion(array []*Score, i, j int) int {
	piovt := array[i]
	left := i + 1
	right := j
	for left != right {
		for left < right && array[left].LessThan(piovt) {
			left++
		}

		for left < right && !array[right].LessThan(piovt) {
			right--
		}

		if left < right {
			array[left], array[right] = array[right], array[left]
		}
	}

	print(array)
	//fmt.Println(i, left)
	if piovt.LessThan(array[left]) {
		left = left - 1
	}

	array[i], array[left] = array[left], array[i]
	return left
}
