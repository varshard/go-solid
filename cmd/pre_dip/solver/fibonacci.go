package solver

import (
	"strconv"
)

func readInput(input string) (int, error) {
	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return n, nil
}

type Fibonacci struct {
	Input int
}

func (f *Fibonacci) SetInput(input string) (err error) {
	f.Input, err = readInput(input)
	return err
}

func (f *Fibonacci) Solve() []string {
	builder := make([]string, 0, f.Input)
	for a, b := 0, 1; a < f.Input; a, b = b, a+b {
		builder = append(builder, strconv.Itoa(b))
	}
	return builder
}
