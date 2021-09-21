package __21

//动态规划：爬楼梯的花费版本
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//滚动数组优化空间复杂度为 n(1)
func minCostClimbingStairs2(cost []int) int {
	n := len(cost)
	pre, cur := 0, 0
	for i := 2; i <= n; i++ {
		pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}

func minCostClimbingStairs3(cost []int) int {

	head := 0
	next := 0
	for i := 2; i <= len(cost); i++ {
		a := next + cost[i-1]
		b := head + cost[i-2]
		if a > b {
			head, next = next, b
		} else {
			head, next = next, a
		}
	}
	return next
}
