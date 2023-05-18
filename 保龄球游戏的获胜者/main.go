package main

import "fmt"

func main() {
	//fmt.Println(isWinner([]int{4, 10, 7, 9}, []int{6, 5, 2, 3}))
	fmt.Println(isWinner([]int{10, 2, 2, 3}, []int{3, 8, 4, 5}))
}

func isWinner(player1 []int, player2 []int) int {
	sum1, sum2 := 0, 0

	sum1 = cal(player1)
	sum2 = cal(player2)

	fmt.Println(sum1, sum2)

	if sum1 > sum2 {
		return 1
	}

	if sum1 == sum2 {
		return 0
	}

	return 2
}

func cal(player []int) int {
	sum := 0
	for i := 0; i < len(player); i++ {
		sum += player[i]
		if i >= 2 && (player[i-1] == 10 || player[i-2] == 10) {
			sum += player[i]
		} else if i >= 1 && player[i-1] == 10 {
			sum += player[i]
		}
	}
	return sum
}
