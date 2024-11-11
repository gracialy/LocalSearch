package main

import (
	"time"
)

type Stochastic struct {
	Experiment
	MaxIterations int
}

func NewStochastic(cube *Cube, maxIterations int) *Stochastic {
	s := &Stochastic{
		MaxIterations: maxIterations,
	}
	s.Experiment = *NewExperiment(cube)
	return s
}

func (s *Stochastic) Run() {
	start := time.Now()
	init := s.Experiment.GetState(0)
	current := init.Clone()
	neighbor := current.Clone()

	for i := 0; i < s.MaxIterations; i++ {
		neighbor.Copy(current)
		neighbor.FindRandomNeighbor()
		if neighbor.Value < current.Value {
			current.Copy(neighbor)
		}

		s.Experiment.AppendState(current)
	}

	s.Experiment.SetRuntime(time.Since(start))
}
