package main

import "strategy/pkg"

var (
	start      = 10
	end        = 100
	straregies = []pkg.Strategy{
		&pkg.PublicTransportStrategy{},
		&pkg.RoadStrategy{},
		&pkg.WalkStrategy{},
	}
)

func main() {
	nav := pkg.Navigator{}
	for _, strategy := range straregies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}
}

/*
Паттерн "Стратегия" (Strategy) позволяет изменять алгоритмы, используемые объектом, без изменения самого объекта.

Метод SetStrategy позволяет навигатору динамически изменять стратегию маршрутизации, не изменяя сам класс навигатора.
*/