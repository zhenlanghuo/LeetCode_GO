package main

import "fmt"

func main() {
	//fmt.Println(minOperations(54))
	fmt.Println(minOperations(1))
}

func minOperations(n int) int {
	arr := make([]int, 0)
	arr = append(arr, 0)
	for n > 0 {
		arr = append(arr, n%2)
		n = n / 2
	}

	//fmt.Println(arr)
	//arr = append([]int{})
	arr = append(arr, 0)
	arr = append(arr, 0)
	fmt.Println(arr)
	pre1Count := 0
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			//fmt.Println(i, "1", pre1Count)
			pre1Count++
		} else {
			//fmt.Println(i, "0", pre1Count)
			if pre1Count > 1 {
				count++
				arr[i] = 1
				pre1Count = 1
			} else {
				count += pre1Count
				pre1Count = 0
			}
		}
		//fmt.Println(i, arr)
	}

	return count
}
