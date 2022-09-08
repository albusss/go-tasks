package cyclicSort

func CyclicSort() {

	in := []int{2, 3, 1, 8, 2, 3, 5, 1}
	i := 0

	for i < len(in) {
		j := in[i] - 1 //  находим место где оно должно быть

		if in[j] != in[i] { // если число было на своем месте, то тут оба элеменна ссылаются на одно и тоже число
			temp := in[j]
			in[j] = in[i]
			in[i] = temp
		} else {
			i++
		}
	}
}
