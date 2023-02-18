package main

import "fmt"

func main() {
	fmt.Println(makeIntegerBeautiful(16, 6))
	fmt.Println(makeIntegerBeautiful(94598, 6))
}

func makeIntegerBeautiful(n int64, target int) int64 {
	arr := make([]int, 0)
	sum := int64(0)
	for n > 0 {
		arr = append(arr, int(n%10))
		sum += n % 10
		n = n / 10
	}
	fmt.Println(arr)

	ans := int64(0)
	r := 0
	temp := 1
	for r < len(arr) && sum > int64(target) {
		//fmt.Println(r, arr, temp, ans)
		if arr[r] == 0 {
			temp *= 10
			r++
			continue
		}
		ans += int64((10 - arr[r]) * temp)
		temp *= 10
		sum -= int64(arr[r] - 1)
		r++

		for r < len(arr) && arr[r] == 9 {
			//fmt.Println(r, arr, temp, ans, 111)
			temp *= 10
			sum -= 9
			r++
		}

		if r < len(arr) {
			arr[r]++
		}
	}

	return ans
}
