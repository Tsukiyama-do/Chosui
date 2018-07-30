package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	primes2 := [5]int{ 50, 70, 110, 92, 22}

	var s []int = primes[1:4]
	var s2 []int = primes2[1:4]
	s = append(s, s2...)
	fmt.Println(s)
}
