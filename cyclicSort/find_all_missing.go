package cyclicSort

import "fmt"

func FindAllMissing() {
	in := []int{2, 3, 1, 8, 2, 3, 5, 1}

	n := len(in)
	i := 0
	var res []int

	for i < n {
		itsPlace := in[i] - 1

		if in[i] != in[itsPlace] {
			temp := in[itsPlace]
			in[itsPlace] = in[i]
			in[i] = temp
		} else {
			i++
		}
	}

	for i = 0; i < n; i++ {
		if i != in[i]-1 {
			res = append(res, i+1)
		}
	}

	fmt.Println(res)
}
