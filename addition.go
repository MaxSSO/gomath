package gomath

func Addition(args ...int) int {
	total := 0
	for _, arg := range args {
		total += arg
	}
	return total
}
