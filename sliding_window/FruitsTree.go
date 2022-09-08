package sliding_window

func Fruits() int {
	in := []byte{'A', 'B', 'C', 'B', 'B', 'C'}
	mapp := make(map[byte]int)
	start := 0
	//res := make(map[byte]int)
	counter := 0
	max := 0

	for i := 0; i < len(in); i++ {
		right := in[i]
		if _, ok := mapp[right]; !ok {
			mapp[right] = 0
		}
		mapp[right] += 1

		for len(mapp) > 2 {
			left := in[start]
			delete(mapp, left)
			start++
		}

		counter = 0
		for _, v := range mapp {
			counter += v
		}
		if counter > max {
			max = counter
		}
	}

	return max
}
