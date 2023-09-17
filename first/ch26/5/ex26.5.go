package main

import "fmt"

// type Mail struct {
// 	alarm Alarm
// }

// type Alarm struct {
// }

// func (m *Mail) OnRecv() {
// 	m.alarm.Alarm()
// }

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() {
	fmt.Println("rrrrrrrrrrr")
}

func main() {
	mail := &Mail{}

	var listener EventListener = &Alarm{}

	mail.Register(listener)
	mail.OnRecv()
}
