package main

import (
	"fmt"
	"math/big"
	"os"
)

//  printFibonacci - выводит необходимое количество чисел Фибоначчи
func printFibonacci(count int64) {
	var i int64

	f0 := big.NewInt(0)
	f1 := big.NewInt(1)

	fn := big.NewInt(0)

	fmt.Printf("1 число Фибоначчи: %d \n", f0)
	fmt.Printf("2 число Фибоначчи: %d \n", f1)

	for i = 3; i <= count; i++ {
		fn.Add(f0, f1)
		fmt.Printf("%d число Фибоначчи: %d \n", i, fn)

		f0 = f1
		f1 = fn
	}
}

// isDivisionWithoutRemainder - проверка деления без остатка n на m
func isDivisionWithoutRemainder(n, m int64) bool {
	return n%m == 0
}

// isSimpleNumber - проверяет число на простоту
func isSimpleNumber(number int64) bool {
	var i int64

	for i = 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

//  printSimpleNumbers - выводит необходимое количество простых чисел
func printSimpleNumbers(count int64) {
	var i, printCount int64

	printCount = 0

	for i = 2; printCount < count; i++ {
		if isSimpleNumber(i) {
			printCount++
			fmt.Printf("%d простое число: %d \n", printCount, i)
		}
	}
}

// PrintMenu - вывод меню
func printMenu() int {

	fmt.Println("")
	fmt.Println("Выберите операцию")
	fmt.Println("  1. Вывод на экран чисел Фибоначчи")
	fmt.Println("  2. Проверка на четность")
	fmt.Println("  3. Проверка деления без остатка на 3")
	fmt.Println("  4. Вывод на экран простых чисел")

	fmt.Println("  5. Выход")
	number := getStdinData("Введите номер операции: ")
	fmt.Println("")

	return int(number)
}

// getStdinData - считать данные из ввода
func getStdinData(str string) int64 {
	var data int64

	fmt.Print(str)
	fmt.Fscan(os.Stdin, &data)

	return data
}

// main - точка входа
func main() {

	noExit := true

	for noExit {
		number := printMenu()

		var str string

		switch number {

		case 1:
			// 1. Вывод на экран чисел Фибоначчи
			fmt.Println("Вы выбрали 1")

			str = "Вывод на экран чисел Фибоначчи - введите сколько чисел вывести: "
			printFibonacci(getStdinData(str))

		case 2:
			// 2. Проверка на четность
			fmt.Println("Вы выбрали 2")

			str = "Введите число для проверки на четность: "

			if isDivisionWithoutRemainder(getStdinData(str), 2) {
				fmt.Println("число четное")
			} else {
				fmt.Println("число не четное")
			}

		case 3:
			// 3. Проверка деления без остатка на 3
			fmt.Println("Вы выбрали 3")

			str = "Введите число для проверки деления на 3 без остатка: "

			if isDivisionWithoutRemainder(getStdinData(str), 3) {
				fmt.Println("число делится на 3 без остатка")
			} else {
				fmt.Println("число не делится на 3 без остатка")
			}

		case 4:
			// 4. Вывод на экран простых чисел
			fmt.Println("Вы выбрали 4")

			str = "Вывод на экран простых чисел - введите сколько чисел вывести: "
			printSimpleNumbers(getStdinData(str))

		case 5:
			// 5. Выход
			fmt.Println("Вы выбрали 5 - Выход")
			noExit = false
		}
	}

}
