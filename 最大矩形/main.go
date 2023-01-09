package main

import "fmt"

func main() {
	fmt.Println(maximalRectangle([][]byte{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}))
}

func maximalRectangle(matrix [][]byte) int {
	heightMatrix := make([][]int, len(matrix)+1)
	for i := 0; i < len(heightMatrix); i++ {
		heightMatrix[i] = make([]int, len(matrix[0]))
	}

	for i := 1; i < len(matrix)+1; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i-1][j] == '0' {
				heightMatrix[i][j] = 0
			} else {
				heightMatrix[i][j] = heightMatrix[i-1][j] + 1
			}
		}
	}

	fmt.Println(heightMatrix)

	max := 0
	for i := 1; i < len(matrix)+1; i++ {
		temp := largestRectangleArea(heightMatrix[i])
		if max < temp {
			max = temp
		}
	}

	return max
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	max := 0
	// 单调栈, 保存的是下标
	stack := make([]int, 0)
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		// 当前柱状高度比栈顶的下标对应的柱状高度要小时,
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			// 将栈顶元素出栈，计算以该下标对应的柱状高度为矩形高度的最大矩形面积
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 从topIndex开始左边第一个小于topIndex对应的柱状高度的下标
			leftIndex := stack[len(stack)-1]
			//leftIndex := -1
			//if len(stack) != 0 {
			//	leftIndex = stack[len(stack)-1]
			//}
			// i为topIndex开始右边第一个小于topIndex对应的柱状高度的小标
			rightIndex := i
			if max < (rightIndex-leftIndex-1)*heights[topIndex] {
				max = (rightIndex - leftIndex - 1) * heights[topIndex]
			}
		}

		// 栈为空 或者 当前的柱状高度比栈顶的下标对应的柱状高度还要高时 将当前下标入栈
		stack = append(stack, i)
	}

	return max
}
