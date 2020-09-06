package main

import "fmt"

func main() {

	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))
	fmt.Println(breakingRecords([]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}))

}

func breakingRecords(scores []int32) []int32 {

	var min, max int32 = scores[0], scores[0]
	counts := []int32{0, 0}

	for i := 1; i < len(scores); i++ {
		if scores[i] > max {
			max = scores[i]
			counts[0]++
		} else if scores[i] < min {
			min = scores[i]
			counts[1]++
		}
	}

	return counts
}
