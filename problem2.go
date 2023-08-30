package main

import "fmt"

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func main()  {
	num := 1000
	for i := 1; i < num+1; i++ {
		fmt.Println(fibonacci(2*i))
	}
}