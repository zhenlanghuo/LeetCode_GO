package main

import "fmt"

func main() {
	height := []int{4,2,0,3,2,5}
	fmt.Print(trap(height))
}

func trap(height []int) int {
	leftMax := make([]int, len(height), len(height))
	rightMax := make([]int, len(height), len(height))

	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}

	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > rightMax[i+1] {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	sum := 0
	for i := 0; i < len(height); i++ {
		if !(leftMax[i] > height[i] && rightMax[i] > height[i]) {
			continue
		}
		sum += min(leftMax[i], rightMax[i]) - height[i]
	}

	return sum
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}
