package quadratic

import "testing"

type Equation struct {
	coeff  Coeff
	answer Answer
}

var testspair = []Equation{
	{Coeff{1, -70, 600}, Answer{60, 10, "D > 0, два корня"}},
	{Coeff{9, -30, 25}, Answer{5.0 / 3.0, 0, "D = 0, один корень"}},
	{Coeff{-2, 3, -5}, Answer{0, 0, "D < 0, нет действительных корней"}},
}

func TestQuadraticSet(t *testing.T) {
	for _, pair := range testspair {
		aw := Quadratic(pair.coeff)
		if aw != pair.answer {
			t.Error(
				"For", pair.coeff,
				"expected", pair.answer,
				"got", aw,
			)
		}
	}
}
