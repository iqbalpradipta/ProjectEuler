package main

import "fmt"


func multiple(num int) []int {
	var numbersOfStatic int = 1000
	var slice []int
	for i := 0; i < numbersOfStatic; i++ {
		if i % 3 == 0 || i % 5 == 0{
			num = i
			slice = append(slice, num)
		}
	}
	return slice[:]
}


func main()  {
	num := 0
	fmt.Println(multiple(num))
}