package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Car машина
type Car struct {
	Model string
	Ready int
	Speed int
}

var cars = []*Car{
	&Car{"Ferrari F40", 0, 0},
	&Car{"Toyota Supra", 0, 0},
	&Car{"McLaren F1", 0, 0},
	&Car{"Dodge Charger", 0, 0},
	&Car{"Lada Priora", 0, 0},
}

func main() {
	// Инициализация
	for _, car := range cars {
		car.Ready = 3 + rand.Intn(12)
		car.Speed = 100 + rand.Intn(100)
	}

	// гоночные синхронизаторы
	var ready sync.WaitGroup
	var start sync.WaitGroup
	syncstart := make(chan int, 0)

	// Начало
	fmt.Println("Гонка началась!")

	fmt.Println("Подготовка...")

	for _, car := range cars {
		// увеличиваем счетчик WaitGroup
		ready.Add(1)
		// стартуем горутину
		go func(car *Car) {
			fmt.Printf("Машина %s ждет старта через %d секунд \n", car.Model, car.Ready)
			time.Sleep(time.Duration(car.Ready) * time.Second)
			fmt.Printf(" - Машина %s готова \n", car.Model)

			ready.Done()

			// ждем всех на старте
			start.Add(1)
			<-syncstart

			fmt.Printf("Машина %s будт ехать %d секунд \n", car.Model, 1000/car.Speed)
			time.Sleep(time.Duration(1000/car.Speed) * time.Second)
			fmt.Printf(" - Машина %s финишировала \n", car.Model)

			start.Done()

		}(car)
	}
	// ожидаем когда все подготовятся
	ready.Wait()

	fmt.Println("Общий старт!")
	// даем старт
	close(syncstart)
	start.Wait()

	fmt.Println("Гонка окончена!")

}
