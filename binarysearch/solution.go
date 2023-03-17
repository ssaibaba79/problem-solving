package main

import "fmt"

func main() {
	input := []int{1, 3, 6, 8, 10, 45, 67, 98, 101}
	binarySearch(input, 8)
	binarySearch(input, 67)
	binarySearch(input, 5)

}

func binarySearch(list []int, item int) {

	low := 0
	high := len(list) - 1
	i := 0
	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			fmt.Printf("Found %d from %d in %d iterations \n", item, len(list), i)
			return

		} else if guess < item {
			low = mid + 1

		} else if guess > item {
			high = mid - 1
		}
		i++

	}
	fmt.Printf("After %d iterations, Item %d not found in a list of %d elements  \n", i, item, len(list))
}
