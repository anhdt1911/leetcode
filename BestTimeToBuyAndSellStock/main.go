package main

import "fmt"

func main() {
	prices := []int{7, 1, 4, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	profit, minPrice := 0, prices[0]
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else if v-minPrice > profit {
			profit = v - minPrice
		}
	}
	return profit
}
