package main

import (
	"fmt"
)

func main() {
	// Standard U.S. coin denominations in cents
	denominations := []int{1, 5, 10, 25, 50}

	// Test amounts
	amounts := []int{87, 42, 99, 33, 7}

	for _, amount := range amounts {
		// Find minimum number of coins
		minCoins := MinCoins(amount, denominations)

		// Find coin combination
		coinCombo := CoinCombination(amount, denominations)

		// Print results
		fmt.Printf("Amount: %d cents\n", amount)
		fmt.Printf("Minimum coins needed: %d\n", minCoins)
		fmt.Printf("Coin combination: %v\n", coinCombo)
		fmt.Println("---------------------------")
	}
}

// MinCoins returns the minimum number of coins needed to make the given amount.
// If the amount cannot be made with the given denominations, return -1.
func MinCoins(amount int, denominations []int) int {
    
    if amount == 0{
        return 0
    }
	
	coins := []int{}

	for i:=0;i<=amount;i++{
		coins = append(coins, i)	
	}
	dp := make([]int, len(coins))

	for i:=1;i<len(dp);i++{
		dp[i]=amount+1
		for j:=0;j<len(denominations);j++{
			if denominations[j] == i{
				dp[i]=1
			}
		}
	}

	for i := 1; i <= amount; i++ {
    	for _, coin := range denominations {
			if coin <= i {
				candidate := dp[i-coin] + 1

				if candidate < dp[i] {
					dp[i] = candidate
				}
        	}
		
    }

}
if dp[amount] == amount+1 {
    return -1
}




	return dp[amount]
}

// CoinCombination returns a map with the specific combination of coins that gives
// the minimum number. The keys are coin denominations and values are the number of
// coins used for each denomination.
// If the amount cannot be made with the given denominations, return an empty map.
 
func CoinCombination(amount int, denominations []int) map[int]int {
	result := make(map[int]int)

	if amount == 0 {
		return result
	}

	// dp[i] = minimum number of coins needed to make i
	dp := make([]int, amount+1)

	// usedCoin[i] = coin that was used to obtain the best solution for i
	usedCoin := make([]int, amount+1)

	// Initialize dp
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	// Calculate minimum coins
	for i := 1; i <= amount; i++ {
		for _, coin := range denominations {
			if coin <= i {
				candidate := dp[i-coin] + 1

				if candidate < dp[i] {
					dp[i] = candidate
					usedCoin[i] = coin
				}
			}
		}
	}

	// Impossible to make the amount
	if dp[amount] == amount+1 {
		return result
	}

	// Reconstruct the combination
	current := amount

	for current > 0 {
		coin := usedCoin[current]

		result[coin]++

		current -= coin
	}

	return result
}