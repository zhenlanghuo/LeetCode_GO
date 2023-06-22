package main

import "fmt"

func main() {
	fmt.Println(distanceTraveled(9, 2))
}

func distanceTraveled(mainTank int, additionalTank int) int {
	ans := 0

	for mainTank != 0 {
		if mainTank >= 5 {
			ans += 5 * 10
			mainTank -= 5
			if additionalTank >= 1 {
				mainTank++
				additionalTank--
			}
		} else {
			ans += mainTank * 10
			mainTank = 0
		}
	}

	return ans
}
