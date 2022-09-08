package TwoPointers

import "fmt"

func Squaring() {
	in := []int{-6, -4, 0, 2, 2}
	result := make([]int, len(in))
	highest := len(in) - 1
	squarLeft := 0
	squarRight := 0
	left := 0
	right := highest

	for left < right {
		squarLeft = in[left] * in[left]
		squarRight = in[right] * in[right]

		if squarLeft > squarRight {
			result[highest] = squarLeft
			left++
		} else {
			result[highest] = squarRight
			right--
		}
		highest -= 1
	}

	fmt.Println("Squaring: ", result)
}
