package main

import "fmt"

type command interface {
	execute()
}

type device interface {
	on()
	off()
}

type onCommand struct {
	device device
}

func (r *onCommand) setDevice(device device) {
	r.device = device
}

func (loc *onCommand) execute() {
	loc.device.on()
}

type offCommand struct {
	device device
}

func (r *offCommand) setDevice(device device) {
	r.device = device
}

func (loc *offCommand) execute() {
	loc.device.off()
}

type lightBulb struct {}

func (lb lightBulb) on() {
	fmt.Println("Light bulb ON!")
}

func (lb lightBulb) off() {
	fmt.Println("Light bulb OFF!")
}

type remote struct {
	cmd command
}

func (r *remote) setCommand(cmd command) {
	r.cmd = cmd
}

func (r *remote) pressButton() {
	r.cmd.execute()
}

func main() {
	lightBulb := &lightBulb{}

	onCommand := &onCommand{}
	onCommand.setDevice(lightBulb)
	
	offCommand := &offCommand{}
	offCommand.setDevice(lightBulb)

	oncmd := &remote{}
	oncmd.setCommand(onCommand)
	oncmd.pressButton()

	offcmd := &remote{}
	offcmd.setCommand(offCommand)
	offcmd.pressButton()
}

/*
Паттерн Command (Команда) используется для инкапсуляции запросов в виде объектов,
позволяя вам параметризовать объекты с операциями, задерживать выполнение операций или поддерживать отмену выполненных операций

1. Каждый запрос на включение или выключение лампочки инкапсулирован в отдельный объект команды (onCommand и offCommand).
Эти команды реализуют интерфейс command, который определяет метод execute.

2. Разделение ответственности: Объекты команды (onCommand и offCommand) знают только о том, как выполнить свои действия,
но не знают, кто или когда их будет вызывать. Лампочка (lightBulb) знает, как выполнять свои действия (on и off),
но не знает, когда эти действия будут запрашиваться.

3. Инвокер: Объект remote выступает в роли инвокера, который хранит команду и вызывает её метод execute.
Он не знает, что конкретно делает команда, он просто вызывает её.
*/

