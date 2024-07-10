package pkg

import "time"

// карты
type Card struct {
	Name string
	Balance float64
	Bank *Bank
}

// реализация поведения, когда карта является инициатором запроса в банк
func (card Card) CheckBalance() error {
	println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 800)
	return card.Bank.CheckBalance(card.Name)
}