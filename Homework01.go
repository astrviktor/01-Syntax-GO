package main

import (
	"fmt"
	"math"
	"os"
)

// ConvertRUBtoUSD - конвертирование рублей в доллары
func ConvertRUBtoUSD(Money float64) {
	const CostUSD = 63.74

	MoneyUSD := Money / CostUSD
	fmt.Printf("%.2f RUB = %.2f USD \n", Money, MoneyUSD)
}

// SquareAndOtherPifagor - площадь, периметр, гипотенуза прямоугольного треугольника
func SquareAndOtherPifagor(A, B float64) {
	C := math.Sqrt(A*A + B*B)
	P := A + B + C
	S := A * B / 2
	fmt.Println("Прямоугольный треугольник: ")
	fmt.Printf("  Стороны A = %.2f, B = %.2f, C = %.2f \n", A, B, C)
	fmt.Printf("  Площадь S = %.2f \n", S)
	fmt.Printf("  Периметр P = %.2f \n", P)
}

// DepositSum5Years - подсчет суммы вклада через 5 лет
func DepositSum5Years(Money, Percent float64) {
	Sum := Money * (1.0 + Percent/100.0)
	Sum = Sum * (1.0 + Percent/100.0)
	Sum = Sum * (1.0 + Percent/100.0)
	Sum = Sum * (1.0 + Percent/100.0)
	Sum = Sum * (1.0 + Percent/100.0)

	fmt.Printf("%.2f RUB, %.2f %% на 5 лет = %.2f RUB \n", Money, Percent, Sum)
}

// PrintMenu - вывод меню
func PrintMenu() int {

	fmt.Println("")
	fmt.Println("Выберите операцию")
	fmt.Println("  1. Конвертация рублей в доллары")
	fmt.Println("  2. Площадь, периметр, гипотенуза прямоугольного треугольника")
	fmt.Println("  3. Сумма вклада через 5 лет")
	fmt.Println("  4. Выход")
	Number := GetStdinData("Введите номер операции: ")
	fmt.Println("")

	return int(Number)
}

// GetStdinData - считать данные из ввода
func GetStdinData(Str string) float64 {
	var Data float64

	fmt.Print(Str)
	fmt.Fscan(os.Stdin, &Data)

	return Data
}

// main - точка входа
func main() {

	NoExit := true

	for NoExit {
		Number := PrintMenu()

		switch Number {
		case 1:
			// 1. Конвертация рублей в доллары
			fmt.Println("Вы выбрали 1")
			ConvertRUBtoUSD(GetStdinData("Введите сумму в рублях: "))
		case 2:
			// 2. Площадь, периметр, гипотенуза прямоугольного треугольника
			fmt.Println("Вы выбрали 2")
			SquareAndOtherPifagor(GetStdinData("Введите сторону А: "), GetStdinData("Введите сторону B: "))
		case 3:
			// 3. Сумма вклада через 5 лет
			fmt.Println("Вы выбрали 3")
			DepositSum5Years(GetStdinData("Введите сумму: "), GetStdinData("Введите процент: "))
		case 4:
			// Выход
			fmt.Println("Вы выбрали 4 - Выход")
			NoExit = false
		}
	}

}
