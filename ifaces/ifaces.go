package main

import (
	"fmt"
)

// Payer - интерфейс платильщика
type Payer interface {
	Pay(int) error
}

// -----------------

// Wallet - кошелек
type Wallet struct {
	Cash int
}

// Pay - метод удовлетворяющий интерфейсу
func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}

	w.Cash -= amount
	return nil
}

// -----------------

// Card - банковская карта
type Card struct {
	Balance int
	Number  string
}

// Pay - метод удовлетворяющий интерфейсу
func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Не хватает денег на карте")
	}

	c.Balance -= amount
	return nil
}

// -----------------

// Buy - функция покупки
func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func main() {

	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	myCard := &Card{Balance: 50, Number: "0000 0000 0000 0000 0000"}
	Buy(myCard)
}
