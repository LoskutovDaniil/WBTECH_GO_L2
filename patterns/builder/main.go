package main

import "builder/pkg"

// создание завода по производству компьютеров asus и hp
func main() {
	asusCollector := pkg.GetCollector("asus")
	hpCollector := pkg.GetCollector("hp")

	factory := pkg.NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()
	
	factory.SetCollector((asusCollector))
	pc := factory.CreateComputer()
	pc.Print()
}

// Паттерн "Строитель" позволяет поэтапно строить сложные объекты, обеспечивая их создание через последовательность шагов (методов).
// Разделение процесса создания объекта: Процесс создания объекта Computer разделён на отдельные шаги (установка ядер, памяти и т.д.),
// которые реализуются в конкретных строителях (AsusCollector и HpCollector).
// Управление последовательностью создания: Фабрика (Factory) управляет последовательностью шагов для создания объекта,
// что упрощает контроль над процессом.
// Гибкость конфигурации: Вы можете легко менять конфигурации создаваемых объектов,
// подставляя разные реализации Collector (например, asusCollector и hpCollector), что делает код гибким и расширяемым.



