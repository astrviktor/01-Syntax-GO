package main

import (
	"fmt"
	"time"
)

//func main() {
//	go spinner(50 * time.Millisecond)
//	const n = 42
//	fibN := fibonacci(n)
//	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
//}

func main() {
	go spinner(50 * time.Millisecond)

	// вращаем 10 секунд
	time.Sleep(10 * time.Second)
	fmt.Print("\rBye\n")
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
