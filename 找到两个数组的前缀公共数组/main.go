package main

func main() {

}

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	count := make([]int, n+1)
	ans := make([]int, n+1)
	for i := 0; i < n; i++ {
		ans[i+1] = ans[i]
		count[A[i]]++
		if count[A[i]] == 2 {
			ans[i+1]++
		}
		count[B[i]]++
		if count[B[i]] == 2 {
			ans[i+1]++
		}
	}
	return ans[1:]
}
