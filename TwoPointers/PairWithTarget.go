package TwoPointers

func F1() []int {
	in := []int{2, 5, 9, 11}
	target := 15
	curSum := 0
	start := 0
	end := len(in) - 1

	for start < end {
		curSum = in[start] + in[end]

		if curSum == target {
			return []int{start, end}
		}

		if curSum > target {
			end--
		}
		if curSum < target {
			start++
		}
	}

	return []int{}
}
