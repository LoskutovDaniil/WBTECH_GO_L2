package main

import "fmt"

// Интерфейс Handler
type Handler interface {
	SetNext(Handler)
	Handle(string)
}

// Конкретный обработчик
type ConcreteHandler struct {
	next Handler
	name string
}

// Устанавливает следующий обработчик в цепочке
func (h *ConcreteHandler) SetNext(next Handler) {
	h.next = next
}

// Обрабатывает запрос и передает его следующему обработчику, если он есть
func (h *ConcreteHandler) Handle(s string) {
	fmt.Println(h.name + " handled the request.")
	if h.next != nil {
		h.next.Handle(s)
	}
}

func main() {
	// Создание обработчиков
	handler1 := &ConcreteHandler{name: "Handler 1"}
	handler2 := &ConcreteHandler{name: "Handler 2"}
	handler3 := &ConcreteHandler{name: "Handler 3"}

	// Установка следующего обработчика в цепочке
	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	// Запуск обработки запроса
	handler1.Handle("Hello")
}

/*
Паттерн "Цепочка ответственности" (Chain of Responsibility) позволяет передавать запрос по цепочке потенциальных обработчиков,
пока один из них не обработает запрос.

1. Создание обработчиков: Создаются три конкретных обработчика handler1, handler2 и handler3.
2. Установка цепочки: Устанавливается цепочка обработчиков: handler1 -> handler2 -> handler3.
3. Запрос "Hello" передается первому обработчику в цепочке (handler1), который обрабатывает запрос
и передает его дальше по цепочке, если это необходимо.
*/