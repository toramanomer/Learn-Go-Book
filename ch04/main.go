package main

import (
	"fmt"
	"io"
	"math/rand/v2"
)

func exercise1() []int {
	randInts := make([]int, 100)

	for i := range 100 {
		randInts[i] = rand.IntN(101)
	}

	return randInts
}

func exercise2(w io.Writer, ints []int) {
	for _, randInt := range ints {
		switch by2, by3 := randInt%2 == 0, randInt%3 == 0; {
		case by2 && by3:
			fmt.Fprint(w, "Six!")
		case by2:
			fmt.Fprint(w, "Two!")
		case randInt%3 == 0:
			fmt.Fprint(w, "Three!")
		default:
			fmt.Fprint(w, "Never mind")
		}
	}
}

func main() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i // Shadowing
		fmt.Println(total)
	}
	fmt.Println("main:", total)
}
