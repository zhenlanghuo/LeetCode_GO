package main

func main() {

}

func differenceOfSum(nums []int) int {
	sum1 := 0
	sum2 := 0

	for _, num := range nums {
		sum1 += num
		for num > 0 {
			sum2 += num % 10
			num = num / 10
		}
	}

	return abs(sum1, sum2)
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
