package maxProfit
// Dynamic programming
func maxProfit(prices []int) int {
	length := len(prices)
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = -prices[0]
	hold, sell := 0, 0
	for i := 1; i < length; i++ {
		sell, hold = dp[0], dp[1]
		if sell > hold+prices[i] {
			dp[0] = sell
		} else {
			dp[0] = hold + prices[i]
		}
		if hold > sell-prices[i] {
			dp[1] = hold
		} else {
			dp[1] = sell - prices[i]
		}
	}
	return dp[0]
}
