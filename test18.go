package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	primes2 := [5]int{ 50, 70, 110, 92, 22}

	var s []int = primes[1:4]
	var s2 []int = primes2[1:4]
	s = append(s, s2...)
	fmt.Println(s)

	j := 0;
    for  i := 0  ; i < 8 ; i++ {
		fmt.Printf("This is %d.  My god.\n", i * 5)
		j =  j + i*5
	}

	fmt.Printf("Total is %d.  Sorry.\n", j)
	sm := 0
	for {
		if sm > 50 {
			fmt.Println("Loop end")
			break
		}
		sm++
	}
	fmt.Println("Loop exited.")
}
