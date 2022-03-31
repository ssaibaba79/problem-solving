package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	/**
	12345

	n=1234, r=5
	n=123, r=54
	n=12, r=543
	n=1, r=5432
	n=0, r=54321

	54321
	*/

	r := 0
	for n != 0 {
		r = r * 10
		r = r + n%10
		n = n / 10
		fmt.Printf("n=%d, r=%d \n", n, r)
	}

	fmt.Println(r)

}
