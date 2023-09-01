package main

import "fmt"

func primeNumbers(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i ++ {
		if num % i == 0{
			return false
		}
	}
	return true
}

func findPrime(n int) int {
	if n <= 0 {
		return 0
	}

	c := 0
	num := 2

	for{
		if primeNumbers(num) {
			c++
		}
		if c == n {
			return num
		}
		num++
	}
}


func main()  {
	n := 10001
	fmt.Println(findPrime(n))
}