package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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
