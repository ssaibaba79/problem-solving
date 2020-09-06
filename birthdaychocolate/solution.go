package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(birthday([]int32{1, 2, 2, 1, 3, 1}, 4, 2))
	fmt.Println(birthday([]int32{3, 2, 3, 2, 4, 1}, 5, 2))
	fmt.Println(birthday([]int32{1, 2, 2, 1, 3, 1}, 5, 3))
	fmt.Println(birthday([]int32{1, 2, 2, 1, 3, 1}, 10, 6))
	fmt.Println(birthday([]int32{1, 2, 2, 1, 3, 1}, 3, 1))
	fmt.Println(birthday([]int32{1, 2, 2, 1, 3, 1}, 8, 7))
	fmt.Println(birthday([]int32{1, 2, 2, 2, 2, 1}, 4, 2))
}

func birthday(s []int32, d int32, m int32) int32 {

	start := time.Now()
	var nWays int32 = 0
	n := int32(len(s))

	// no need to check.  block > n
	if m > n {
		fmt.Printf("Result : s=%v, m=%v, d=%v, ways=%v\n", s, m, d, nWays)
		timeTest(start)
		return 0
	}

	// block  = n
	if m == n {
		if checkSum(s, d) {
			nWays = 1
		}

		fmt.Printf("Result : s=%v, m=%v, d=%v, ways=%v\n", s, m, d, nWays)
		timeTest(start)
		return nWays
	}

	if m == 1 {
		f := make(map[string]int)
		for _, v := range s {
			if v == d {
				k := strconv.Itoa(int(v))
				_, found := f[k]
				if !found {
					f[k] = 1
					nWays++
				}
			}
		}
		fmt.Printf("Result : s=%v, m=%v, d=%v, ways=%v\n", s, m, d, nWays)
		timeTest(start)
		return nWays
	}

	fmt.Printf("Test Inputs: s=%v,d=%v, m=%v\n", s, d, m)
	fmt.Println("---------------------------------------")
	patterns := make(map[string]int)
	var i int32 = 0
	for i = 0; i <= n-m; i++ {
		b := make([]int32, m)
		copy(b, s[i:i+m])

		if checkSum(b, d) {
			pk := fmt.Sprint(b)
			_, present := patterns[pk]
			if !present {
				patterns[pk] = 1
				nWays++
			}
		}

		fmt.Printf("%v checksum: %v,d=%v, ways=%v\n", i, b, d, nWays)
	}

	fmt.Printf("Result : s=%v, m=%v, d=%v, ways=%v\n", s, m, d, nWays)
	fmt.Println("---------------------------------------")
	timeTest(start)
	return nWays

}

func timeTest(start time.Time) {
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Completed in : %v seconds\n", elapsed)
}

func checkSum(b []int32, d int32) bool {
	var sum int32 = 0
	for _, v := range b {
		sum += v
	}

	if sum == d {
		return true
	}

	return false

}
