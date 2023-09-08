package main

import "fmt"

type Attacker interface {
	Name() string
	Attack(DamageTaker, int)
}

type DamageTaker interface {
	DealDamage(Attacker, int)
}

type Player struct {
	name string
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Attack(dt DamageTaker, damage int) {
	dt.DealDamage(p, damage)
}

type Monster struct {
	hp int
}

func (m *Monster) DealDamage(attacker Attacker, damage int) {
	m.hp -= damage
	if m.hp < 0 {
		fmt.Printf("killed by %s", attacker.Name())
	}
}

func main() {

	var t DamageTaker = &Monster{50}
	var a Attacker = &Player{"Player A"}

	a.Attack(t, 100)
}
