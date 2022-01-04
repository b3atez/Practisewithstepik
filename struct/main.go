package main

import "fmt"

type Player struct {
	On    bool
	Ammo  int
	Power int
}

func (p *Player) Shoot() bool {
	if p.On != false {
		if p.Ammo >= 1 {
			p.Ammo -= 1
			return true
		}
	}

	return false
}

func (p *Player) RideBike() bool {
	if p.On != false {
		if p.Power >= 1 {
			p.Power -= 1
			return true
		}
	}

	return false
}

func main() {
	testStrct := &Player{
		On:    true,
		Ammo:  20,
		Power: 30,
	}
	fmt.Println(testStrct.RideBike())
}
