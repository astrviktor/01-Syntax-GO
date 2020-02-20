package main

import "fmt"

//марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель, открыты
//ли окна, насколько заполнен объем багажника

// tLuggage - тип багажник
type tLuggage struct {
	volume int // объем
	use    int // использовано
}

// tCar - тип машина
type tCar struct {
	model         string // марка
	year          int    // год выпуска
	isEngineStart bool   // запущен ли двигатель
	isWindowsOpen bool   // открыты ли окна
}

// tTruck - тип грузовик
type tTruck struct {
	car     tCar     // грузовик это машина
	luggage tLuggage // у грузовика есть кузов
}

// main - точка входа
func main() {

	car1 := tCar{"Ferrari F40", 1990, true, true}

	car2 := tCar{"Ford Focus", 2005, false, false}

	truck1 := tTruck{tCar{"Toyota Tundra", 2016, true, false}, tLuggage{1000, 100}}

	truck2 := tTruck{
		tCar{
			model:         "Tesla Cybertruck",
			year:          2022,
			isEngineStart: true,
			isWindowsOpen: true,
		},
		tLuggage{2000, 2000},
	}

	car2.model = "Ford Fiero"
	car2.year = 1991

	fmt.Println(car1)
	fmt.Println(truck1)

	fmt.Printf("%+v \n", car2)
	fmt.Printf("%+v  \n", truck2)

}
