package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("---------------------------------------")
	start := time.Now()
	fmt.Println(divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))
	elapsed(start)
	fmt.Println("---------------------------------------")

	fmt.Println("---------------------------------------")
	start = time.Now()
	fmt.Println(divisibleSumPairs(6, 5, []int32{1, 3, 2, 6, 1, 2}))
	elapsed(start)
	fmt.Println("---------------------------------------")
}

func elapsed(start time.Time) {
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Completed in : %v seconds\n", elapsed)
}

// Complete the divisibleSumPairs function below.
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var i, count int32 = 0, 0

	//fmt.Printf("Test Inputs: n=%v,k=%v, ar=%v\n", n, k, ar)
	//fmt.Println("---------------------------------------")

	for i = 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				//fmt.Printf("Sum ar[%v] + ar[%v] of pair [%v,%v] is divisible by %v\n", i, j, ar[i], ar[j], k)
				count++
			}
		}
	}

	return count

}
