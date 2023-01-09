package main

import (
	"fmt"
)

func main() {
	fmt.Println(videoStitching([][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 3))
}

func videoStitching(clips [][]int, time int) int {

	// maxn[i]记录在clips中, 左端点为i的最大右端点
	maxn := make([]int, time)
	for _, clip := range clips {
		if clip[0] < time && clip[1] > maxn[clip[0]] {
			maxn[clip[0]] = clip[1]
		}
	}

	// last记录目前能到达的最大的右端点, prev记录上一个被使用的区间的结束位置
	last, prev := 0, 0
	count := 0
	for i, v := range maxn {
		// 表示检查区间[i, i+1]
		if v > last {
			last = v
		}
		// i==last, 表示目前最大只能到last, 区间[i, i+1]缺失
		if i == last {
			return -1
		}
		// i等于prev说明上一个区间已经不满足当前区间了, 需要使用新的区间
		if i == prev {
			count++
			// 新的区间的结束位置为last
			prev = last
		}
	}

	return count
}

//func videoStitching(clips [][]int, time int) int {
//	// dp[i]代表剪辑i秒视频需要的最少片段
//	dp := make([]int, time+1)
//	for i := 0; i <= time; i++ {
//		dp[i] = math.MaxInt32
//	}
//	dp[0] = 0
//
//	for i := 1; i <= time; i++ {
//		for _, clip := range clips {
//			if clip[0] < i && clip[1] >= i {
//				dp[i] = min(dp[i], dp[clip[0]]+1)
//			}
//		}
//	}
//
//	if dp[time] == math.MaxInt32 {
//		return -1
//	}
//
//	return dp[time]
//}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func videoStitching(clips [][]int, time int) int {
//	if time == 0 {
//		return 0
//	}
//
//	sort.Slice(clips, func(i, j int) bool {
//		return clips[i][0] < clips[j][0]
//	})
//
//	if clips[0][0] != 0 {
//		return -1
//	}
//	maxEnd := 0
//	//candidate := clips[0]
//	right := 0
//	count := 0
//	for right < len(clips) {
//		nextEnd := 0
//		// 基于上一个已选的片段的end, 选择下一个选择的片段，下一个选择的片段的start需要小于等于上一个已选片段的end
//		for right < len(clips) && clips[right][0] <= maxEnd {
//			// 下一个选择的片段的end越到越好
//			if clips[right][1] > nextEnd {
//				nextEnd = clips[right][1]
//			}
//			right++
//		}
//		if maxEnd == nextEnd {
//			return -1
//		}
//		count++
//		maxEnd = nextEnd
//		if nextEnd >= time {
//			return count
//		}
//	}
//
//	return -1
//}
