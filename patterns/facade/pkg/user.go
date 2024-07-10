package pkg

// пользователь
type User struct {
	Name string
	Card *Card 
}

// реализация поведения, когда пользователь хочет узнать баланс
func (user User) GetBalance() float64 {
	return user.Card.Balance
}