package main

import (
	"fmt"
	"io"
	"math"
)

func exercise1() {
	i := 20
	f := float64(i)
	fmt.Println(i)
	fmt.Printf("%.2f\n", f)
}

func exercise2(w io.Writer) {
	const value = 10
	var i = value
	var f float64 = value

	fmt.Fprintln(w, i)
	fmt.Fprintln(w, f)
}
func exercise3() (byte, int32, uint64) {
	var (
		b      byte   = math.MaxUint8
		smallI int32  = math.MaxInt32
		bigI   uint64 = math.MaxUint64
	)

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)

	return b, smallI, bigI
}

func main() {
	exercise3()
}
