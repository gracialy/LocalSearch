package main

import (
	"math"
	"math/rand"
	"time"
)

const (
	SA_MAX             = 100000000000
	COOLING_RATE       = 0.999999
	CAP_BOLTZMANN_PLOT = CAP_PLOT
	CAP_T              = 0.1
)

type SimulatedAnnealing struct {
	Experiment
	T               float64
	InitialT        float64
	Boltzmann       []float64
	stuck           int
	ActualIteration int
}

func NewSimulatedAnnealing(cube *Cube, initialT float64) *SimulatedAnnealing {
	sa := &SimulatedAnnealing{}
	sa.Experiment = *NewExperiment(cube)
	sa.Boltzmann = make([]float64, 0)
	sa.stuck = 0
	sa.InitialT = initialT
	return sa
}

func (sa *SimulatedAnnealing) Run() {
	start := time.Now()

	init := sa.Experiment.GetState(0)
	current := init.Clone()
	neighbor := current.Clone()

	for i := 1; i < SA_MAX; i++ {
		sa.schedule(i)
		sa.ActualIteration = i

		if sa.T <= CAP_T {
			// very close to 0 since the T will never touch 0
			break
		}

		neighbor.Copy(current)
		neighbor.FindRandomNeighbor()
		delta := neighbor.Value - current.Value
		probability := sa.probability(float64(delta))
		random := rand.Float64()
		sa.AppendProbability(probability)

		if delta <= 0 {
			current.Copy(neighbor)
			sa.Boltzmann[len(sa.Boltzmann)-1] = sa.InitialT
		} else if probability > random {
			current.Copy(neighbor)
			sa.stuck++
		}

		sa.Experiment.AppendState(current)
	}

	sa.Experiment.SetRuntime(time.Since(start))
}

func (sa *SimulatedAnnealing) schedule(t int) {
	sa.T = sa.InitialT * math.Pow(COOLING_RATE, float64(t))
}

func (sa *SimulatedAnnealing) probability(delta float64) float64 {
	return math.Exp(-delta / sa.T)
}

func (sa *SimulatedAnnealing) AppendProbability(probability float64) {
	sa.Boltzmann = append(sa.Boltzmann, probability)
}
