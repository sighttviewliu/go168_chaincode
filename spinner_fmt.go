package main
import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45 
	fibN := fib(n)
	fmt.Printf("\nFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) { //spin show
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\n%c", r)
			time.Sleep(delay)
		}
	}
}


func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

