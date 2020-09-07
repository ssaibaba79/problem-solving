package main

import (
	"fmt"
	"time"
)

// Problem
//https://www.hackerrank.com/challenges/migratory-birds/problem

func migratoryBirds(arr []int32) int32 {

	//fmt.Printf("Test Inputs: arr=%v\n", arr)

	sightingCounts := make(map[int32]int32)

	for _, v := range arr {
		sightingCounts[v] = sightingCounts[v] + 1
	}

	//fmt.Println(sightingCounts)

	countsMap := make(map[int32][]int32)
	var max int32 = 0
	for birdType, count := range sightingCounts {
		//fmt.Printf("birdType=%v, count=%v\n", birdType, count)
		if count > max {
			max = count
		}
		birdtypes, found := countsMap[count]
		if !found {
			birdtypes = []int32{}

		}
		birdtypes = append(birdtypes, birdType)
		countsMap[count] = birdtypes

	}
	//fmt.Println(countsMap)
	maxSightings := countsMap[max]
	var t int32 = 0
	if len(maxSightings) > 0 {
		t = maxSightings[0]
		for i := 1; i < len(maxSightings); i++ {
			if maxSightings[i] < t {
				t = maxSightings[i]
			}
		}
	}

	return t

}

func main() {

	fmt.Println("---------------------------------------")
	start := time.Now()
	fmt.Println(migratoryBirds([]int32{1, 3, 2, 5, 1, 2}))
	elapsed(start)
	fmt.Println("---------------------------------------")

	fmt.Println("---------------------------------------")
	start = time.Now()
	fmt.Println(migratoryBirds([]int32{2, 3, 2, 4, 3, 2, 3}))
	elapsed(start)
	fmt.Println("---------------------------------------")
}

func elapsed(start time.Time) {
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Completed in : %v seconds\n", elapsed)
}
