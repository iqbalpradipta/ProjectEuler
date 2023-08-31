package main

import (
	"fmt"
	"strconv"
)


func palindromeFunc(num int) int {
	if num < 0 {
		return -1
	}

	numStr := strconv.Itoa(num)
	palindromeStr := ""

	for i := len(numStr)-1; i >= 0; i-- {
		palindromeStr += string(numStr[i])
	}

	palindrome, _ := strconv.Atoi(palindromeStr)
	return palindrome
}

func main()  {
	var dumb int
	for i := 100; i < 1000; i++ {
		dumb = i * i
		if dumb == palindromeFunc(dumb) {
			fmt.Printf("Perkalian %v X %v adalah palindrome dengan hasil = %d\n",i, i, dumb)
		}
	}
}