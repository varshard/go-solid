package solver

import "strconv"

type FizzBuzz struct {
	Input int
}

func (f *FizzBuzz) SetInput(input string) (err error) {
	f.Input, err = readInput(input)
	return err
}

func (f *FizzBuzz) Solve() []string {
	builder := make([]string, 0)
	for i := 1; i <= f.Input; i++ {
		if i%15 == 0 {
			builder = append(builder, "FizzBuzz")
		} else if i%3 == 0 {
			builder = append(builder, "Fizz")
		} else if i%5 == 0 {
			builder = append(builder, "Buzz")
		} else {
			builder = append(builder, strconv.Itoa(i))
		}
	}
	return builder
}
