package main

import (
	"fmt"
	"log"
	"state/pkg"
)

func main() {
	vendingMachine := pkg.NewVendingMachine(1, 10)
	// запросил можно ли купить товар
	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// добавил деньги
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// произошла выдача товара
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// дополнили двумя товарами машину
	fmt.Println()
	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	// пользователь запросил можно ли купить
	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	// вставил деньги
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// автомат выдал товар покупателю
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

/*
Паттерн "Состояние" (State) позволяет объекту изменять свое поведение в зависимости от его состояния.

Когда методы RequestItem, InsertMoney, DispenseItem или AddItem вызываются на торговом автомате, они делегируются текущему состоянию,
которое определяет, как должен вести себя торговый автомат в этом состоянии. В зависимости от выполнения этих методов,
торговый автомат может переходить в другое состояние, изменяя свое поведение.
*/