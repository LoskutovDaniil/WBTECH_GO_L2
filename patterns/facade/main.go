// реализация поведения когда приходит пользователь и пытается купить товар

package main

import (
	"facade/pkg"
	"fmt"
)

var (
	bank = pkg.Bank{
		Name:  "БАНК",
		Cards: []pkg.Card{},
	}

	card1 = pkg.Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}

	card2 = pkg.Card{
		Name:    "CRD-2",
		Balance: 5,
		Bank:    &bank,
	}

	user = pkg.User{
		Name: "Покупатель-1",
		Card: &card1,
	}

	user2 = pkg.User{
		Name: "Покупатель-2",
		Card: &card2,
	}

	prod = pkg.Product{
		Name:  "Сыр",
		Price: 150,
	}

	shop = pkg.Shop{
		Name: "SHOP",
		Products: []pkg.Product{
			prod,
		},
	}
)

func main() {
	println("[Банк] Выпуск карты")
	bank.Cards = append(bank.Cards, card1, card2)

	fmt.Printf("[%s]\n", user.Name)
	// Ф-ия Sell Является фасадом над бизнес логикой по безналичному расчету
	err := shop.Sell(user, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}

	fmt.Printf("[%s]\n", user2.Name)
	// Ф-ия Sell Является фасадом над бизнес логикой по безналичному расчету
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
}

/* Функция Sell является фасадом, потому что она упрощает взаимодействие
с более сложной системой и скрывает детали реализации бизнес-логики по безналичному расчету от клиента (в данном случае, вызывающего кода).
Фасадный паттерн предоставляет унифицированный интерфейс для взаимодействия с подсистемой,
тем самым уменьшая сложность и облегчая использование системы. */
