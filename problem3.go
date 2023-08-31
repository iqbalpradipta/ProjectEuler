package main

import "fmt"


func primeNumber(l int) []int {
	var num int
	var slice []int
	for i := 1; i <= l; i++ {
		if l % i == 0 {
			num = i
			slice = append(slice, num)
		}
	}
	return slice[:]
}

func main()  {
	l := 600851475143
	fmt.Println(primeNumber(l))
}