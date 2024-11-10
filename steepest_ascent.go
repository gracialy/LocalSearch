package main

import (
	"time"
)

type SteepestAscent struct {
	Experiment
	ActualIteration int
}

func NewSteepestAscent(c *Cube) *SteepestAscent {
	sta := &SteepestAscent{}
	sta.Experiment = *NewExperiment(c.Clone())
	sta.ActualIteration = 0
	return sta
}

func (sta *SteepestAscent) Run() {
	start := time.Now()
	init := sta.Experiment.GetState(0)
	current := init.Clone()
	i := 0

	for {
		neighbor := current.FindBestNeighbor()
		if neighbor.Value >= current.Value {
			break
		}
		current.Copy(neighbor)

		i++
		sta.ActualIteration = i
		sta.Experiment.AppendState(current)
	}

	sta.Experiment.SetRuntime(time.Since(start))
}
