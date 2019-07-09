package sort

import (
	"fmt"
)

//CoinChange 硬币找零的动态规划
func CoinChange(values, coinUsed []int, money int)  {

	for cents := 1; cents <= money; cents++ {
		minCoins := cents  // 默认采用最小 1 元的硬币 -- 1 元的情况下，最小硬币数为 1。
		
		for kind := 0; kind < len(values); kind++ {
			if values[kind] <= cents {
				// 硬币组合为从大到小
				// 如果当前遍历的硬币比当前找零的钱数小，那么最后一位必定是该硬币。然后最小硬币数为当前找零的钱数减去当前遍历的硬币值对应的最小硬币数加 1
				temp := coinUsed[cents-values[kind]] + 1
				if temp < minCoins {
					minCoins = temp
				}
			}
		}
		coinUsed[cents] = minCoins
		fmt.Println("")
	}
}