package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	num := readNum()

	output := execute(os.Args[1], num)
	display(output)
}

func display(output []string) {
	fmt.Print(strings.Join(output, "\n"))
}

func readNum() int {
	arg := os.Args[2]
	num, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}
	return num
}

func execute(name string, num int) []string {
	switch name {
	case "fibonacci":
		return Fibonacci(num)
	case "fizzbuzz":
		return FizzBuzz(num)
	default:
		panic("unsupported game")
	}
}

func Fibonacci(num int) []string {
	builder := make([]string, 0, num)
	for a, b := 0, 1; a < num; a, b = b, a+b {
		builder = append(builder, fmt.Sprintf("%d", b))
	}
	return builder
}
func FizzBuzz(num int) []string {
	builder := make([]string, 0, num)
	for i := 1; i <= num; i++ {
		if i%15 == 0 {
			builder = append(builder, "FizzBuzz")
		} else if i%3 == 0 {
			builder = append(builder, "Fizz")
		} else if i%5 == 0 {
			builder = append(builder, "Buzz")
		} else {
			builder = append(builder, fmt.Sprintf("%d", i))
		}
	}
	return builder
}
