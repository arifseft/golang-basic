package main

import (
	"fmt"
	"log"

	"strings"
)

func Fib(n int, a int, b int, arr []int) []int {
	if b == 0 {
		b = 1
	}
	if n > 0 {
		arr = append(arr, a)
		return Fib(n - 1, b, a + b, arr)
	}

	return arr
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func main() {
	var num int
	var err error
	var arr []int

	fmt.Print("Enter value of num: ")
	_, err = fmt.Scanf("%d", &num)
	if err != nil {
		log.Fatal(err.Error())
	}

	fib := Fib(num, 0, 1, arr)
	fmt.Printf("Fibonacci: %s", arrayToString(fib, ", "))
	fmt.Println()
}
