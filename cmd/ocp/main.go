package main

import (
	"fmt"
	"os"
	"solid/cmd/ocp/solver"
	"strings"
)

var solvers map[string]solver.Solver

func main() {
	game := getSolver(os.Args[1])
	display(solve(game, os.Args[2]))
}

func display(output []string) {
	fmt.Print(strings.Join(output, "\n"))
}

func getSolver(name string) solver.Solver {
	if s, ok := solvers[name]; ok {
		return s
	}
	panic("unsupported solver")
}

func solve(q solver.Solver, input string) []string {
	if err := q.SetInput(input); err != nil {
		panic(err)
	}
	return q.Solve()
}

func init() {
	solvers = map[string]solver.Solver{
		"fibonacci": &solver.Fibonacci{},
		"fizzbuzz":  &solver.FizzBuzz{},
	}
}
