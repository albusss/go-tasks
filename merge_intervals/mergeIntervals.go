package merge_intervals

import (
	"fmt"
	"math"
	"sort"
)

func MergeIntervals() {
	in := [][]int{{2, 5}, {1, 4}, {7, 9}}

	sort.Slice(in, func(i, j int) bool {
		if in[i][0]-in[j][0] < 0 {
			return true
		} else {
			return false
		}
	})
	intervals := make([][]int, 0)

	start := in[0][0]
	end := in[0][1]

	for i := 1; i < len(in); i++ {
		currentInterval := in[i]

		if currentInterval[0] < end {
			max := math.Max(float64(currentInterval[1]), float64(end))
			end = int(max)
		} else {
			intervals = append(intervals, []int{start, end})
			start = currentInterval[0]
			end = currentInterval[1]
		}
	}
	intervals = append(intervals, []int{start, end})

	fmt.Println(intervals)
}
