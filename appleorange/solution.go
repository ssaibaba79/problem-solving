package main

import "fmt"

func main() {

	fmt.Println("Test1:")
	countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6, 8, -7})
	fmt.Println("Test2:")
	countApplesAndOranges(7, 11, 7, 15, []int32{-2, 2, 1}, []int32{5, -6, 7, -3})

}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {

	counts := []int32{0, 0}

	counts[0] = countFruits(s, t, a, apples)
	counts[1] = countFruits(s, t, b, oranges)

	for _, count := range counts {
		fmt.Println(count)
	}

}

func countFruits(s int32, t int32, treePos int32, fruits []int32) int32 {
	var count int32 = 0
	for _, d := range fruits {

		var p int32 = treePos + d
		if p >= s && p <= t {
			count++
		}

	}
	return count
}
