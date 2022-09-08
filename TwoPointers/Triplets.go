package TwoPointers

import (
	"fmt"
	"sort"
)

func Triplets() {
	in := []int{-3, 0, 1, 2, -1, 1, -2}
	sort.Ints(in)

	for i := 0; i < len(in); i++ {
		if i > 0 && in[i] == in[i-1] {
			continue
		}

		find(in, -in[i], i+1)
	}
}

func find(in []int, targetSum, left int) {
	right := len(in) - 1
	currentSum := 0
	for left < right {
		currentSum = in[left] + in[right]
		if currentSum == targetSum {
			fmt.Println([]int{-targetSum, in[left], in[right]})
			left++
			right--
			for left < right && in[left] == in[left-1] {
				left++
			}
			for left < right && in[right] == in[right+1] {
				right--
			}
		} else if currentSum < targetSum {
			left++
		} else if currentSum > targetSum {
			right--
		}
	}
}
