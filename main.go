package main

import "fmt"

var m = make(map[int]int, 0)

func main() {
	for _, n := range []int{1, 2, 9, 98, 6} {
		x := fib(n)
		fmt.Println(n, "fib", x)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	var f int
	if v, ok := m[n]; ok {
		f = v
	} else {
		f = fib(n-2) + fib(n-1)
		m[n] = f
	}

	return f
}

// func main() {
// 	iPanic(5)
// }

// func iPanic(i int) {
// 	if i > 0 {
// 		iPanic(i - 1)
// 	}
// 	panic("I'm outta here")
// }
