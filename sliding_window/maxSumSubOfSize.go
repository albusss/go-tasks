package sliding_window

func Max() int {
	in := []int{2, 1, 5, 1, 3, 2}
	k := 3
	max := 0
	sum := 0
	start := 0

	for i := start; i < len(in); i++ {
		sum = sum + in[i]

		if i >= k-1 {
			if sum > max {
				max = sum
			}
			sum -= in[start]
			start++
		}
	}
	return max
}
