package puzzles

func cons(a, b int) func(func(int, int) int) int {
	pair := func(f func(a, b int) int) int {
		return f(a, b)
	}
	return pair
}

func car(pair func(func(int, int) int) int) int {
	return pair(func(a, b int) int {
		return a
	})
}

func cdr(pair func(func(int, int) int) int) int {
	return pair(func(a, b int) int {
		return b
	})
}
