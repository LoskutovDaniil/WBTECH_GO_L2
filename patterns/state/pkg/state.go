package pkg

// общий интерфейс состояния
type State interface {
	AddItem(int) error						// добавление товара
	RequestItem() error						// запрос интересующего нас товара
	InsertMoney(money int) error			// внесение денег
	DispenseItem() error					// выдача предмета
}

