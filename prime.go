package main

import (
	"fmt"
	"log"

	"strings"
)

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func Prime(n int) []int {
	var prime []int
	var i int
	var count int
	var c int

	if n >= 1 {
		prime = append(prime, 2)
	}

	i = 3
	for count = 2; count <= n; i++ {
		for c = 2; c < i; c++ {
			if i%c == 0 {
				break
			}
		}
		if c == i {
			prime = append(prime, i)
			count++
		}
	}

	return prime
}

func main() {
	var num int
	var err error

	fmt.Print("Enter value of num: ")
	_, err = fmt.Scanf("%d", &num)
	if err != nil {
		log.Fatal(err.Error())
	}

	prime := Prime(num)
	fmt.Printf("Prime: %s", arrayToString(prime, ", "))
	fmt.Println()
}
