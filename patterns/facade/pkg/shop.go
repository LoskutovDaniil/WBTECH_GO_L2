package pkg

import (
	"errors"
	"fmt"
	"time"
)

// магазин
type Shop struct {
	Name     string
	Products []Product
}

// реализация поведение магазина, когда к нему приходит пользователь и пытается купить какой - нибудь товар
// Является фасадом над бизнес логикой по безналичному расчету
func (shop Shop) Sell(user User, product string) error {
	println("[Магазин] Запрос к пользователю, для получения остатка на карте")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка - может ли [%s] купить товар! \n", user.Name)
	time.Sleep(time.Millisecond * 500)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недорстаточно средств для покупки товара!")
		}
		fmt.Printf("[Магазин] Товар [%s] - куплен!\n", prod.Name)
	}
	return nil
}
