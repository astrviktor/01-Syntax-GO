package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Point - точка
type Point struct {
	x int
	y int
}

// Sum - сумма двух точек, результат в саму переменную
func (p *Point) Sum(p1, p2 Point) {
	p.x = p1.x + p2.x
	p.y = p1.y + p2.y
}

// IsCorrectPoint - проверка точки на корректность
func (p *Point) IsCorrectPoint() bool {
	if p.x < 1 || p.y < 1 {
		return false
	}
	if p.x > 8 || p.y > 8 {
		return false
	}

	return true
}

// ToString - в формат шахматной доски
func (p *Point) ToString() string {

	var convertorToString = map[int]string{
		1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "F", 7: "G", 8: "H"}

	return convertorToString[p.x] + strconv.Itoa(p.y)
}

// Для конвертации из шахматного формата в Point
var convertorToPoint = map[string]int{
	"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8,
	"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8}

// Ходы коня
var horsePoints = [8]Point{
	{1, 2}, {-1, 2}, {2, 1}, {-2, 1},
	{1, -2}, {-1, -2}, {2, -1}, {-2, -1}}

// ConvertStrToPoint - строку вида "E2" в (6,2)
func ConvertStrToPoint(str string) (Point, error) {
	var p = Point{0, 0}
	var err error

	if len(str) != 2 {
		err = errors.New("Некорректная точка")
		return p, err
	}

	str = strings.ToUpper(str)

	p.x = convertorToPoint[string(str[0])]
	p.y = convertorToPoint[string(str[1])]

	if p.IsCorrectPoint() {
		return p, nil
	}

	err = errors.New("Некорректная точка")
	return p, err
}

func main() {
	fmt.Println("Введите стартовую позицию коня (например E2)")
	input := ""

	fmt.Scanln(&input)

	if horsePoint, err := ConvertStrToPoint(input); err == nil {

		fmt.Println("Доступны следующие ходы")

		var point Point

		for idx := 0; idx < 8; idx++ {
			point.Sum(horsePoint, horsePoints[idx])
			if point.IsCorrectPoint() {
				fmt.Println(point.ToString())
			}
		}

	} else {
		fmt.Println("Возникла ошибка:", err)
	}

}
