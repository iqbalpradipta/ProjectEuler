package main

import "fmt"

func sumOfSquare(a int) int {
	var num int
	for i := 1; i <= a; i++ {
		num += i*i
	}
	return num
}

func squareOfSum(b int) int {
	var result int
	var sum int
	for i := 1; i <= b; i++ {
		sum += i
		result = sum*sum
	}
	return result
}

func main()  {
	var difference int
	a := 100
	b := 100
	difference = squareOfSum(b) - sumOfSquare(a)
	fmt.Printf("Sum of Square first %v = %d\n",a, sumOfSquare(a))
	fmt.Printf("Square of Sum first %v = %d\n",b, squareOfSum(b))
	fmt.Println("====================================================================")
	fmt.Printf("Hence the difference between the sum of the squares of the first ten \nnatural numbers and the square of the sum is\n%v - %v = %d\n",squareOfSum(b),sumOfSquare(a) ,difference)
}