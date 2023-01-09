package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{30,60,90}))
}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		result[i] = 0
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if temperatures[i] > temperatures[top] {
				result[top] = i - top
				stack = stack[:len(stack)-1]
				continue
			}
			break
		}
		stack = append(stack, i)
	}

	return result
}
