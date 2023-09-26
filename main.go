package main

import (
	"errors"
	"fmt"
)

type Jumper interface {
	WhereAmI() int      // выводит текущее положение кузнечика на линейке
	Jump() (int, error) // кузнечик прыгает к зерну. Выводит новое положение кузнечика, или ошибку, если он уже ест зерно
}

type Grasshopper struct {
	position int
	target   int
}

func (g *Grasshopper) WhereAmI() int {
	return g.position
}

func (g *Grasshopper) Jump() (int, error) {
	if g.position == g.target {
		return g.position, errors.New("error")
	}

	maxJump := 5
	distanceToTarget := g.target - g.position

	if distanceToTarget <= maxJump {
		g.position = g.target
		return g.target, nil
	}
	jumpDistance := 1

	if jumpDistance < 1 || jumpDistance > maxJump {
		return g.position, errors.New("error")
	}

	g.position += jumpDistance
	return g.position, nil
}

func PlaceJumper(place, target int) Jumper {
	return &Grasshopper{
		position: place,
		target:   target,
	}
}

const (
	place  = 0
	target = 3
)

func main() {
	g := PlaceJumper(place, target)
	fmt.Println(g.WhereAmI())
	for {
		currPlace, err := g.Jump()
		if err != nil {
			break
		}
		fmt.Println(currPlace)
	}
}
