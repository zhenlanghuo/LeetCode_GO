package main

func main() {

}

func countTriplets(nums []int) int {
	countMap := make([]int, 1<<16)
	for _, num1 := range nums {
		for _, num2 := range nums {
			countMap[num1&num2]++
		}
	}

	ans := 0
	for _, num1 := range nums {
		for num2, count := range countMap {
			if num1&num2 == 0 {
				ans += count
			}
		}
	}

	return ans
}
