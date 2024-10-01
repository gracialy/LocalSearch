package main

import (
	"math/rand"
	"time"
)

const (
	NMAX = 5000000
)

type Stochastic struct {
	Cube
}

func NewStochastic(cube *Cube) *Stochastic {
	st := &Stochastic{
		Cube: *cube,
	}

	return st
}

func (st *Stochastic) Run() {
	current := &Stochastic{Cube: *st.Cube.Copy()}

	for i := 0; i < NMAX; i++ {
		neighbor := &Stochastic{Cube: *current.Cube.Copy()}
		neighbor.Random()
		if neighbor.Cube.Value > current.Cube.Value {
			current = neighbor
		}
	}
	st.Cube = current.Cube
}

func (st *Stochastic) Random() {
	rand.Seed(time.Now().UnixNano())
	x1 := uint8(rand.Intn(5))
	y1 := uint8(rand.Intn(5))
	z1 := uint8(rand.Intn(5))
	x2, y2, z2 := x1, y1, z1

	for x1 == x2 && y1 == y2 && z1 == z2 {
		x2 = uint8(rand.Intn(5))
		y2 = uint8(rand.Intn(5))
		z2 = uint8(rand.Intn(5))
	}

	st.Swap(x1, y1, z1, x2, y2, z2)

	st.Cube.EvalValue()
}
