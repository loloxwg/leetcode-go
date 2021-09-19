package __15

//股票买入时间2
//贪心算法
func MaxProfit2(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += Max(0, prices[i]-prices[i-1])
	}
	return
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
