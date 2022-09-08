package TwoPointers

func Duplecates() int {
	in := []int{2, 3, 3, 3, 6, 9, 9}

	next := 1
	cur := 0

	for next < len(in) {
		if in[cur] == in[next] {
			next++
		}
		if in[cur] != in[next] {
			in[cur+1] = in[next]
			cur++
			next++
		}
	}

	return cur + 1
}
