package array

func MaxProfit(prices []int) int {
	m := 0
	buy := 0
	sell := 1
	for sell < len(prices) {
		if prices[buy] < prices[sell] {
			p := prices[sell] - prices[buy]
			if p > m {
				m = p
			}
			sell++
		} else {
			buy = sell
			sell++
		}
	}
	return m
}
