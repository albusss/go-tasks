package Dynamic

import "fmt"

func KnapStack() {
	profits := []int{1, 6, 10, 16}
	weights := []int{1, 2, 3, 5}
	capacity := 7

	dp := make(map[int]map[int]int)
	res := knapRecur(profits, weights, capacity, 0, dp)

	fmt.Println(res)
}

func knapRecur(profits, weights []int, cap, curIndex int, dp map[int]map[int]int) int {

	if cap <= 0 || curIndex >= len(profits) {
		return 0
	}

	//---
	if _, ok := dp[curIndex]; !ok {
		dp[curIndex] = make(map[int]int)
	}
	if _, ok := dp[curIndex][cap]; ok {
		return dp[curIndex][cap]
	}
	//---

	profit1 := 0

	if weights[curIndex] <= cap {
		profit1 = profits[curIndex] + knapRecur(profits, weights, cap-weights[curIndex], curIndex+1, dp)
	}

	profit2 := knapRecur(profits, weights, cap, curIndex+1, dp)

	if profit2 > profit1 {
		dp[curIndex][cap] = profit2

		return dp[curIndex][cap]
	}

	dp[curIndex][cap] = profit1

	return dp[curIndex][cap]
}
