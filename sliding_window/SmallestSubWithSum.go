package sliding_window

func F() int {
	in := []int{3, 4, 1, 1, 6}
	s := 8
	sum := 0
	start := 0
	min := len(in)

	for i := 0; i < len(in); i++ {
		sum += in[i]

		for sum >= s {
			curMin := i - start + 1
			if curMin < min {
				min = curMin
			}

			sum -= in[start]
			start++
		}
	}

	if min == len(in) {
		return 0
	}

	return min
}
