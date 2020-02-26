package main

// Нужно поставить переменну $GOPATH или скопировать в src
// папку calculator

import (
	"calculator"
	"fmt"
)

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}
		if input == "exit" {
			break
		}
		if input == "help" {
			fmt.Println("Справка: вводите математические операции")
			continue
		}
		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
