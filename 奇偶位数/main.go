package main

func main() {

}

func evenOddBit(n int) []int {
	ans := []int{0, 0}
	for i := 0; n > 0; i++ {
		if n%2 == 1 {
			ans[i%2]++
		}
		n = n / 2
	}
	return ans
}
