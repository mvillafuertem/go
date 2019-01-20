package main

import (
	"fmt"
	"time"
)

func main() {

	go animacion(1 * time.Millisecond)
	i := 45
	n := fibonacci(i)
	fmt.Printf("\rFibonacci(%d) = %d\n", i, n)

}

func animacion(t time.Duration) {
	a := []string{"....", "...", "..", "."}
	for {
		for _, c := range a {
			fmt.Printf("\r%s", c)
			time.Sleep(t)
		}
	}
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
