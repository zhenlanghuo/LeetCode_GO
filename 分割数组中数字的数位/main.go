package main

func main() {

}

func separateDigits(nums []int) []int {
	result := make([]int, 0)
	temp := make([]int, 0)
	for _, num := range nums {
		for num > 0 {
			temp = append(temp, num%10)
			num = num / 10
		}
		reverse(temp)
		result = append(result, temp...)
		temp = temp[:0]
	}
	return result
}

func reverse(temp []int) {
	l, r := 0, len(temp)-1
	for l < r {
		temp[l], temp[r] = temp[r], temp[l]
		l++
		r--
	}
}
