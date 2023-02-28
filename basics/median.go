package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(median([]float64{3, 1, 2}))    // 2
	fmt.Println(median([]float64{3, 1, 4, 2})) // 2.5
}

func median(nums []float64) float64 {
	// work on a copy, don't change the input
	vals := make([]float64, len(nums))
	copy(vals, nums)
	sort.Float64s(vals)

	i := len(vals) / 2
	if len(vals)%2 == 1 { // odd number of values
		return vals[i]
	}

	return (vals[i-1] + vals[i]) / 2
}
