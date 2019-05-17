package main

import (
	"fmt"
	"log"
)

func Sum(x int, y int) int {
	return x + y
}

func main() {
	var num1 int
	var num2 int
	var err error

	fmt.Print("Enter value of num1: ")
	_, err = fmt.Scanf("%d", &num1)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print("Enter value of num2: ")
	_, err = fmt.Scanf("%d", &num2)
	if err != nil {
		log.Fatal(err.Error())
	}

	sum := Sum(num1, num2)
	fmt.Printf("Sum: %d", sum)
	fmt.Println()
}