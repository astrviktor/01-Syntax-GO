package quadratic

import "math"

// Coeff - коэффициенты квадратного уравнения
type Coeff struct {
	a, b, c float64
}

// Answer - решение квадратного уравнения
type Answer struct {
	x1, x2 float64
	msg    string
}

// Quadratic - подсчет корней квадратного уравнения
func Quadratic(coeff Coeff) (answer Answer) {
	a := coeff.a
	b := coeff.b
	c := coeff.c

	var x1, x2 float64
	var msg string

	// Дискриминант
	D := b*b - 4*a*c

	// Дискриминант < 0
	if D < 0 {
		msg = "D < 0, нет действительных корней"

		return Answer{0, 0, msg}
	}

	if D == 0 {
		msg = "D = 0, один корень"
		x1 = -b / (2 * a)
		x2 = 0
		return Answer{x1, x2, msg}
	}

	msg = "D > 0, два корня"
	x1 = (-b + math.Sqrt(D)) / 2 * a
	x2 = (-b - math.Sqrt(D)) / 2 * a

	return Answer{x1, x2, msg}
}
