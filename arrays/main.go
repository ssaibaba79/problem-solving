package main

import (
	"fmt"
	"unsafe"
)

/*
  Arrays and Slices
*/
func main() {

	// standard declaration
	// var <name> [length]type
	var arr [5]int

	// var <name> [length]type =[length]type {comma seperated values}
	var months [12]string = [12]string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}
	var arr1 [3]int = [3]int{1, 2, 3}

	fmt.Println(arr)
	// shortcut initialization
	arr2 := [4]int{1, 2, 3, 4}

	// large array using range [0,0,0,0,0, 100, 5, -1]
	codes := [...]int{5: 100, 5, -1}

	// Not allowed as arr1 and arr2 are of different lengths and considered as
	// different types. ( unique to Go)
	//if arr1 == arr2 {}

	// iterate through array
	fmt.Println("arr1:")
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("arr1[%v] %p %v %d\n", i, &arr1[i], arr1[i], unsafe.Sizeof(arr1[i]))
	}

	// iterate through an array using range operator
	fmt.Println("arr2:")
	for i, v := range arr2 {
		fmt.Printf("arr2[%v]=%v \n", i, v)
	}

	// 2D array
	var matrix [4][4]int = [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 9}, {1, 1, 1, 1}, {0, -1, 4, 4}}
	printSlice(matrix[:])

	// N-dimensional array
	var polytype [][][][]int = [][][][]int{
		{{{0, 0}, {1, 1}}, {{0, 0}, {1, 1}}}, {{{0, 1}, {1, 1}}}}
	printSlice(polytype)

	// array to slice [:]
	printSlice(months[:])
	printSlice(codes[:])

	// slices
	// slice from an Array
	summer := months[5:8]
	printSlice(summer)

	var s []int = []int{1, 2, 3, 4}
	printSlice(s)

	// using make
	var s2 = make([]string, 10)
	printSlice(s2)

	s2 = append(s2, "a", "b", "c")
	printSlice(s2)

	word := []int{1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(word)
	printSlice(reverse(word[:]))

}

func reverse(slice []int) []int {

	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// Generic print method
func printSlice[T any](arr []T) {
	fmt.Printf("%p len=%d, cap-%d %v \n", arr, len(arr), cap(arr), arr)
}
