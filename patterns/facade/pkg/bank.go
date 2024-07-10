package pkg

import (
	"errors"
	"fmt"
	"time"
)

// банк
type Bank struct {
	Name string
	Cards []Card
}

// реализация поведения банка, когда банку приходит запрос по номеру карты для проверки баланса
func (bank Bank) CheckBalance(cardNumber string) error {
	println(fmt.Sprintf("[Банк] Получение остатка по карте %s", cardNumber))
	time.Sleep(time.Millisecond * 300)
	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств!")
		}
	}
	println("[Банк] Остаток положительный!")
	return nil
}