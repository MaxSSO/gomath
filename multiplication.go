package gomath

func Multiplication(args ...int) int {
	total := 1
	for _, arg := range args {
		total *= arg
	}
	return total
}
