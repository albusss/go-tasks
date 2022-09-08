package sliding_window

import "fmt"

func LongestSubWithKDist() {
	in := "araaci"
	k := 2
	start := 0
	keys := 0
	max := 0
	chMap := make(map[byte]int)

	for i := 0; i < len(in); i++ {
		char := in[i]
		if _, ok := chMap[char]; !ok {
			chMap[char] = 0
			keys++
		}
		chMap[char]++

		for keys > k {
			startChar := in[start]
			chMap[startChar]--
			if chMap[startChar] == 0 {
				keys--
				delete(chMap, startChar)
			}
			start++
		}

		l := i - start + 1
		if l > max {
			max = l
		}
	}

	fmt.Println(max)
}
