package main

import "fmt"

func primeNumbers(l int) []int {
	var num int
	primeNumber100LEGITItsPrimeNumberAfterMuchResearch := []int{2,3,5,7}
	for i := 2; i <= l; i++ {
		if i % 2 != 0 && i % 3 != 0 && i % 5 != 0 && i % 7 != 0{
			num = i
			primeNumber100LEGITItsPrimeNumberAfterMuchResearch = append(primeNumber100LEGITItsPrimeNumberAfterMuchResearch, num)
		}
	}
	return primeNumber100LEGITItsPrimeNumberAfterMuchResearch
}

func main()  {
	prime := 10001
	fmt.Println(primeNumbers(prime))
}