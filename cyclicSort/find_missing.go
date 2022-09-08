package cyclicSort

import "fmt"

func FindMissing() {
	in := []int{4, 0, 3, 1}
	i := 0
	n := len(in)

	for i < n {
		j := in[i]

		if in[i] < n && in[j] != in[i] {
			temp := in[j]
			in[j] = in[i]
			in[i] = temp
		} else {
			i++
		}
	}

	for i = 0; i < n; i++ {
		if in[i] != i {
			fmt.Println(i)
			break
		}
	}

	fmt.Println("Not found")
}
