package main

import (
	"fmt"
)

const maxCount = 10000

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x <= maxCount; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for res := range squares {
		fmt.Println(res)
	}
}
