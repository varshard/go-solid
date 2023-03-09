package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[2]
	num, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}
	switch os.Args[1] {
	case "fibonacci":
		for a, b := 0, 1; a < num; a, b = b, a+b {
			fmt.Printf("%d\n", b)
		}
	case "fizzbuzz":
		for i := 1; i <= num; i++ {
			if i%15 == 0 {
				fmt.Println("FizzBuzz")
			} else if i%3 == 0 {
				fmt.Println("Fizz")
			} else if i%5 == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Printf("%d\n", i)
			}
		}
	}
}
