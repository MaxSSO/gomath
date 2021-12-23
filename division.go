package gomath

func Division(args ...int) int {
	total := 1
	for _, arg := range args {
		total /= arg
	}
	return total
}
