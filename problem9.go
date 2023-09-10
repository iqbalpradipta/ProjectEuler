package main

import (
	"fmt"
	"runtime"
)

/*
	Berdasarkan soal rumus pythagoras a<b<c adalah a^2 + b^2 = c^2
	Dik:
		result := 1000
*/

func main() {
	runtime.GOMAXPROCS(5)
	result := 1000
	var hasil int
	for i := 0; i < result; i++ {
		hasilA := i * i
		for j := 0; j < result; j++ {
			hasilB := j * j
			for k := 0; k < result; k++ {
				hasilC := k * k
				hasil = hasilA + hasilB + hasilC
				if hasil == result {
					go fmt.Printf("%d^2 + %d^2 + %d^2 = %v\n", i, j, k, result)
				}
			}
		}
	}
}
