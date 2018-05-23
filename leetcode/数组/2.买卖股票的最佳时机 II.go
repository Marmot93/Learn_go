package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		if i+1 == len(prices){
			break
		}
		if prices[i] < prices[i+1] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}

func main() {
	price := []int{2, 65, 7, 7, 7, 8}
	profit := maxProfit(price)
	fmt.Print(profit)
}
