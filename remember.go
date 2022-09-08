package main

import "fmt"

func Remember() {
	in := "oidbcaf"
	pattern := "abc"
	lenth := len(pattern)

	patternmap := make(map[byte]int)
	matched := 0
	start := 0

	for i := 0; i < len(pattern); i++ {
		if _, ok := patternmap[pattern[i]]; !ok {
			patternmap[pattern[i]] = 0
		}
		patternmap[pattern[i]] += 1
	}

	for i := 0; i < len(in); i++ {
		char := in[i]

		if _, ok := patternmap[char]; ok {
			patternmap[char]--
			if patternmap[char] == 0 {
				matched++
			}
		}

		if matched == lenth {
			fmt.Println("DOME")
			return
		}

		if i >= lenth-1 {
			leftChar := in[start]
			start++
			if _, ok := patternmap[leftChar]; ok {
				if patternmap[char] == 0 {
					matched--
				}
				patternmap[char]++

			}
		}

	}
}
