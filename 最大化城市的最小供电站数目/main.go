package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxPower([]int{1, 2, 4, 5, 0}, 1, 2))
	fmt.Println(maxPower([]int{4, 4, 4, 4}, 0, 3))
}

//func maxPower(stations []int, r int, k int) int64 {
//	n := len(stations)
//	sum := make([]int, n+1) // 前缀和
//	for i, x := range stations {
//		sum[i+1] = sum[i] + x
//	}
//	mn := math.MaxInt
//	for i := range stations {
//		stations[i] = sum[min(i+r+1, n)] - sum[max(i-r, 0)] // 电量
//		mn = min(mn, stations[i])
//	}
//	return int64(mn + sort.Search(k, func(minPower int) bool {
//		minPower += mn + 1     // 改为二分最小的不满足要求的值，这样 sort.Search 返回的就是最大的满足要求的值
//		diff := make([]int, n) // 差分数组
//		sumD, need := 0, 0
//		for i, power := range stations {
//			sumD += diff[i] // 累加差分值
//			m := minPower - power - sumD
//			if m > 0 { // 需要 m 个供电站
//				need += m
//				if need > k { // 提前退出这样快一些
//					return true // 不满足要求
//				}
//				sumD += m // 差分更新
//				if i+r*2+1 < n {
//					diff[i+r*2+1] -= m // 差分更新
//				}
//			}
//		}
//		return false
//	}))
//}

func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	stationPreSum := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		stationPreSum[i] = stationPreSum[i-1] + stations[i-1]
	}
	stationSum := make([]int, n)
	minStation := math.MaxInt64

	//fmt.Println(stationPreSum)
	for i := 0; i < n; i++ {
		stationSum[i] = stationPreSum[min(i+r+1, n)] - stationPreSum[max(i-r, 0)]
		minStation = min(stationSum[i], minStation)
	}

	//fmt.Println(stationSum)

	//fmt.Println(minStation)

	check := func(minSta int) int {
		minSta += minStation
		diff := make([]int, n)
		sumD, needSum := 0, 0
		for i := 0; i < n; i++ {
			// 累计差分值
			sumD += diff[i]
			need := minSta - stationSum[i] - sumD
			if need > 0 {
				// 差分更新
				sumD += need
				needSum += need
				if needSum > k {
					return 1
				}
				if i+2*r+1 < n {
					// 差分更新
					diff[i+2*r+1] -= need
				}
			}
		}

		if needSum == k {
			return 0
		}
		return -1
	}

	left := 0
	right := k
	for left <= right {
		mid := (left + right) / 2
		if check(mid) > 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return int64(minStation + left - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
