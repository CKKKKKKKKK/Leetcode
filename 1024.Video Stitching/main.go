package videoStitching

import (
	"math"
)

func videoStitching(clips [][]int, T int) int {
	const inf = math.MaxInt64 - 1
	length := len(clips)
	dp := make([]int, T+1)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 1; i <= T; i++ {
		min := inf
		for j := 0; j < length; j++ {
			if clips[j][0] <= i && clips[j][1] >= i {
				if dp[clips[j][0]] < min {
					min = dp[clips[j][0]]
				}
			}
		}
		dp[i] = min + 1
	}
	if dp[T] == inf+1 {
		return -1
	}
	return dp[T]
}
