package main

import (
	"fmt"
	"math"
)

func main() {
	n := 5
	queries := [][3]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}

	fmt.Println("max:", manipulate(n, queries))
}

func manipulate(n int, queries [][3]int32) int64 {
		output := make([]int32, n+2)
		for i := 0; i < len(queries); i++ {
			a := queries[i][0]
			b := queries[i][1]
			k := queries[i][2]

			output[a] += k
			output[b+1] -= k
		}

		max := getMax(output)
		return max
}

// returns the maximum elements of a slice.
func getMax(arr []int32) int64 {
	var max int64 = -1000
	var sum int64 = 0
	for _, v := range arr {
		sum += int64(v)
		max = int64(math.Max(float64(max), float64(sum)))
	}
	return max
}
