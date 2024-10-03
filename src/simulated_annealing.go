package main

import (
	"math"
	"math/rand"
	"time"
)

const (
	SAMAX = 1000000
	TINIT = 1000.0
	BETA  = 0.833
)

type SimulatedAnnealing struct {
	Stochastic
	ActualIteration int
	T               float64
}

func NewSimulatedAnnealing(cube *Cube) *SimulatedAnnealing {
	st := *NewStochastic(cube)

	sa := &SimulatedAnnealing{
		Stochastic: st,
	}

	sa.ActualIteration = 0
	sa.T = TINIT

	return sa
}

func (sa *SimulatedAnnealing) GetConfiguration() [DIMENSION][DIMENSION][DIMENSION]uint8 {
	return sa.Stochastic.GetConfiguration()
}

func (sa *SimulatedAnnealing) GetValue() int {
	return sa.Stochastic.GetValue()
}

func (sa *SimulatedAnnealing) GetValue1() int {
	return sa.Stochastic.GetValue1()
}

func (sa *SimulatedAnnealing) GetValue2() int {
	return sa.Stochastic.GetValue2()
}

func (sa *SimulatedAnnealing) GetCube() *Cube {
	return sa.Stochastic.GetCube()
}

func (sa *SimulatedAnnealing) GetRuntime() time.Duration {
	return sa.Stochastic.GetRuntime()
}

func (sa *SimulatedAnnealing) GetActualIteration() int {
	return sa.ActualIteration
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
	sa.ActualIteration = 0
	timeStart := time.Now()

	current := sa
	next := current.Clone()

	t := 0
	// for {
	for i := 0; i < SAMAX; i++ {
		t++
		sa.schedule(t)
		if sa.T == 0 {
			break
		}
		sa.ActualIteration++
		next.Copy(current)
		next.Random()
		delta := float64(next.Cube.Value - current.Cube.Value)
		if delta < 0 {
			current.Copy(next)
		} else {
			prob := sa.Boltzmann(delta)
			random := float64(randRange(0, 101)) / 100.0
			// fmt.Printf("t: %d delta: %f T: %e Probability: %f Random: %f Value: %d\n", t, delta, sa.T, prob, random, next.GetValue())
			if prob < random {
				current.Copy(next)
			}
		}
	}

	sa.SetRuntime(time.Since(timeStart))
}

func (sa *SimulatedAnnealing) schedule(t int) {
	sa.T = 1 / (BETA*math.Log(float64(t)+1) + TINIT)
}

func (sa *SimulatedAnnealing) Boltzmann(delta float64) float64 {
	return math.Exp(delta / sa.T)
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
