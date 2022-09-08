package sliding_window

import "fmt"

func Permutation() {
	in := "oidbcafb"
	pattern := "abc"
	matched := 0
	start := 0
	mapNeedle := map[byte]int{}
	keysNum := 0

	for i := 0; i < len(pattern); i++ {
		ch := pattern[i]
		if _, ok := mapNeedle[ch]; !ok {
			mapNeedle[ch] = 0
			keysNum++
		}
		mapNeedle[ch] += 1
	}

	for i := 0; i < len(in); i++ {
		ch := in[i]

		if _, ok := mapNeedle[ch]; ok {
			mapNeedle[ch] -= 1
			if mapNeedle[ch] == 0 {
				matched++
			}
		}
		if matched == keysNum {
			fmt.Println("TRUE")
			return
		}

		if i >= len(pattern)-1 {
			leftCh := in[start]
			start++
			if _, ok := mapNeedle[leftCh]; ok {
				if mapNeedle[leftCh] == 0 {
					matched--
				}
				mapNeedle[leftCh]++
			}
		}
	}

	fmt.Println("FALSE")
	return
}
