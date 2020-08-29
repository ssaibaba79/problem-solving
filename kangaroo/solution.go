package main

import (
	"fmt"
)

func main() {
	fmt.Println(kangaroo(43, 1, 41, 2))
	//fmt.Println(kangaroo(0, 3, 4, 2))

	//fmt.Println(kangaroo(4, 3, 1, 5))
	//fmt.Println(kangaroo(0, 3, 0, 2))
	//fmt.Println(kangaroo(0, 3, 0, 3))

}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

	fmt.Printf("START k1=%v,leap=%v, k2=%v,leap=%v \n", x1, v1, x2, v2)

	if x1 < x1 && v1 < v2 {
		fmt.Printf("STOP Condition : x1 < x1 && v1 < v2 k1=%v,leap=%v, k2=%v,leap=%v \n", x1, v1, x2, v2)
		return "NO"
	}

	if x2 < x1 && v2 < v1 {
		fmt.Printf("STOP Condition : x2 < x1 && v2 < v1  k1=%v,leap=%v, k2=%v,leap=%v \n", x1, v1, x2, v2)
		return "NO"
	}

	if x1 == x2 && v1 != v2 {
		fmt.Printf("STOP Condition : x1 == x2 && v1 != v2  k1=%v,leap=%v, k2=%v,leap=%v \n", x1, v1, x2, v2)
		return "NO"
	}

	if x1 != x2 && v1 == v2 {
		fmt.Printf("STOP Condition : x1 != x2 && v1 == v2  k1=%v,leap=%v, k2=%v,leap=%v \n", x1, v1, x2, v2)
		return "NO"
	}

	k1 := x1
	k2 := x2
	d := abs(k1 - k2)

	var i int32 = 1
	for {
		k1 += v1
		k2 += v2
		dd := abs(k1 - k2)

		fmt.Printf("JUMP %v[k1=%v, k2=%v, delta=%v]\n", i, k1, k2, dd)
		if k1 == k2 {
			fmt.Printf("MEET after JUMP %v at %v[k1=%v, k2=%v, delta=%v]\n", i, k1, k1, k2, dd)
			return "YES"
		}

		if dd >= d {
			fmt.Printf("STOP Condition : Delta not converging. k1=%v, k2=%v,prev_delta=%v,current_delta=%v \n", x1, x2, d, dd)
			break
		}
		i++
	}

	return "NO"
}

func abs(i int32) int32 {
	if i < 0 {
		return i * -1
	}
	return i
}
