package main

func main() {

}

func maxCount(banned []int, n int, maxSum int) int {
	bannedMap := make(map[int]bool)
	for _, num := range banned {
		bannedMap[num] = true
	}

	sum, count := 0, 0
	for i := 1; i <= n; i++ {
		if _, ok := bannedMap[i]; ok {
			continue
		}
		sum += i
		if sum > maxSum {
			break
		}
		count++
	}

	return count
}
