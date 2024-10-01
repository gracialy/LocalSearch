package main

import (
	"math/rand"
	"time"
)

const (
	NMAX = 1000000
)

type Stochastic struct {
	Cube
	Runtime time.Duration
}

func NewStochastic(cube *Cube) *Stochastic {
	st := &Stochastic{
		Cube: *cube,
	}

	st.Runtime = 0

	return st
}

func (st *Stochastic) GetConfiguration() [DIMENSION][DIMENSION][DIMENSION]uint8 {
	return st.Cube.GetConfiguration()
}

func (st *Stochastic) GetValue() uint8 {
	return st.Cube.GetValue()
}

func (st *Stochastic) GetRuntime() time.Duration {
	return st.Runtime
}

func (st *Stochastic) GetCube() *Cube {
	return &st.Cube
}

func (st *Stochastic) SetValue() {
	st.Cube.SetValue()
}

func (st *Stochastic) SetRuntime(runtime time.Duration) {
	st.Runtime = runtime
}

func (st *Stochastic) Clone() *Stochastic {
	return &Stochastic{Cube: *st.Cube.Clone()}
}

func (st *Stochastic) Copy(original *Stochastic) {
	st.Cube.Copy(original.GetCube())
	st.Runtime = original.GetRuntime()
}

func (st *Stochastic) PrintConfiguration() {
	st.Cube.PrintConfiguration()
}

func (st *Stochastic) Run() {
	timeStart := time.Now()

	current := st
	neighbor := current.Clone()

	for i := 0; i < NMAX; i++ {
		neighbor.Copy(current)
		neighbor.Random()
		// fmt.Printf("Neighbor Value: %d, Current Value: %d\n", neighbor.GetValue(), current.GetValue())
		if neighbor.GetValue() > current.GetValue() {
			current.Copy(neighbor)
		}
	}

	st.SetRuntime(time.Since(timeStart))
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
}
