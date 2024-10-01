package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	TMAX = 1000000000
)

type SimulatedAnnealing struct {
	Stochastic
}

func NewSimulatedAnnealing(cube *Cube) *SimulatedAnnealing {
	st := *NewStochastic(cube)

	sa := &SimulatedAnnealing{
		Stochastic: st,
	}

	return sa
}

func (sa *SimulatedAnnealing) GetConfiguration() [DIMENSION][DIMENSION][DIMENSION]uint8 {
	return sa.Stochastic.GetConfiguration()
}

func (sa *SimulatedAnnealing) GetValue() uint8 {
	return sa.Stochastic.GetValue()
}

func (sa *SimulatedAnnealing) GetCube() *Cube {
	return sa.Stochastic.GetCube()
}

func (sa *SimulatedAnnealing) GetRuntime() time.Duration {
	return sa.Stochastic.GetRuntime()
}

func (sa *SimulatedAnnealing) SetValue() {
	sa.Stochastic.SetValue()
}

func (sa *SimulatedAnnealing) SetRuntime(runtime time.Duration) {
	sa.Stochastic.SetRuntime(runtime)
}

func (sa *SimulatedAnnealing) Clone() *SimulatedAnnealing {
	return &SimulatedAnnealing{Stochastic: *sa.Stochastic.Clone()}
}

func (sa *SimulatedAnnealing) Copy(original *SimulatedAnnealing) {
	sa.Stochastic.Copy(&original.Stochastic)
}

func (sa *SimulatedAnnealing) Run() {
	timeStart := time.Now()

	current := sa
	next := current.Clone()

	for t := 1; t < TMAX; t++ {
		T := sa.schedule(t)
		if T == 0 {
			break
		}
		next.Copy(current)
		next.Random()
		delta := float64(next.Cube.Value - current.Cube.Value)
		if delta > 0 {
			current = next
		} else {
			prob := math.Exp(-delta / T)
			random := rand.Float64()
			if prob > random {
				fmt.Printf("Probability: %f Random: %f\n", prob, random)
				current.Copy(next)
			}
		}
	}

	sa.SetRuntime(time.Since(timeStart))
}

func (sa *SimulatedAnnealing) schedule(t int) float64 {
	return 1.0 / math.Log(float64(t)+1)
}
