package puzzles

/*
This problem was asked by Jane Street.

cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and last element of that pair. For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.

Given this implementation of cons:

def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair
Implement car and cdr.
*/
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
