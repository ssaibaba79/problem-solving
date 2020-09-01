package main

import "fmt"

func main() {
	fmt.Println(circularArrayRotation([]int32{1, 2, 3, 4, 5}, 1, []int32{3, 4}))
	fmt.Println(circularArrayRotation([]int32{1, 2, 3, 4, 5}, 2, []int32{3, 4}))
	fmt.Println(circularArrayRotation([]int32{1, 2, 3, 4, 5}, 3, []int32{3, 4}))
	fmt.Println(circularArrayRotation([]int32{1, 2, 3, 4, 5}, 4, []int32{3, 4}))
	fmt.Println(circularArrayRotation([]int32{1, 2, 3, 4, 5}, 101, []int32{3, 4}))
}

// Complete the circularArrayRotation function below.
func bfcircularArrayRotation(a []int32, k int32, queries []int32) []int32 {

	fmt.Printf("Before rotation  : %v\n", a)
	var n int32 = int32(len(a))
	for r := 1; r <= int(k); r++ {
		var i int32 = n - 1
		var t int32 = a[i]
		for i = n - 1; i >= 0; i-- {

			if i == 0 {
				a[i] = t
			} else {
				a[i] = a[i-1]
			}
			//fmt.Printf("a[%v]:%v t=%v\n", i, a[i], t)
		}
		//fmt.Printf("After rotation %v : %v\n", r, a)
	}

	v := []int32{}

	for _, m := range queries {
		v = append(v, a[m])
	}
	return v

}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {

	n := int32(len(a))
	s := k % n
	d := n - s
	t := make([]int32, s)
	copy(t, a[d:])
	//fmt.Printf("%v  %v\n", a, t)

	var i int32 = 0
	for i := n - 1; i >= s; i-- {
		a[i] = a[i-s]
	}
	//fmt.Printf("%v %v\n", a, t)

	for i = 0; i < s; i++ {
		a[i] = t[i]
	}

	//fmt.Printf("%v\n", a)

	v := []int32{}

	for _, m := range queries {
		v = append(v, a[m])
	}
	return v

}
