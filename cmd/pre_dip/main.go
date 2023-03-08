package main

import (
	"fmt"
	"os"
	"solid/cmd/pre_dip/solver"
	"strings"
)

var solvers = Solvers{
	FiboSolver:     &solver.Fibonacci{},
	FizzBuzzSolver: &solver.FizzBuzz{},
}

func main() {
	display(solve(os.Args[1], os.Args[2]))
}

func display(output []string) {
	fmt.Print(strings.Join(output, "\n"))
}

func solve(question string, input string) []string {
	switch question {
	case "fibonacci":
		solvers.FiboSolver.SetInput(input)
		return solvers.FiboSolver.Solve()
	case "fizzbuzz":
		solvers.FizzBuzzSolver.SetInput(input)
		return solvers.FizzBuzzSolver.Solve()
	default:
		panic("unsupported question")
	}
}

type Solvers struct {
	FiboSolver     *solver.Fibonacci
	FizzBuzzSolver *solver.FizzBuzz
}
