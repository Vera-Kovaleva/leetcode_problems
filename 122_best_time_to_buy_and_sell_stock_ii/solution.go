package main

func maxProfit(prices []int) int {
	saved := prices[0]
	profit := 0
	for _, price := range prices {
		if saved < price {
			profit += price - saved
			saved = price
		}
		if price < saved {
			saved = price
		}
	}
	return profit
}

