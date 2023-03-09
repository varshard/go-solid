package main

import (
	"fmt"
	"os"
	"solid/cmd/ocp/solver"
	"strings"
)

var solvers = Solvers{}

func main() {
	game := getSolver(os.Args[1])
	display(solve(game, os.Args[2]))
}

func solve(q solver.Solver, input string) []string {
	// Liskov Substitution, whatever is Solver, I accept
	if err := q.SetInput(input); err != nil {
		panic(err)
	}
	return q.Solve()
}

func display(output []string) {
	fmt.Print(strings.Join(output, "\n"))
}

func getSolver(name string) solver.Solver {
	if s, ok := solvers.Solvers[name]; ok {
		return s
	}
	panic("unsupported question")
}

func init() {
	solvers.Solvers["fibonacci"] = &solver.Fibonacci{}
	solvers.Solvers["fizzbuzz"] = &solver.FizzBuzz{}
}

type Solvers struct {
	// Solver interface allow Solvers to not concern about its solvers
	Solvers map[string]solver.Solver
}
