package main

import (
	"fmt"
	"os"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	CAP_DUMP = 10000
	CAP_PLOT = 10000
)

type Experiment struct {
	State   []Cube
	Runtime time.Duration
}

func NewExperiment(c *Cube) *Experiment {
	experiment := &Experiment{}
	experiment.State = append(make([]Cube, 0), *c.Clone())

	return experiment
}

func (e *Experiment) Dump(name string) {
	file, err := os.Create("txt/" + name + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	limit := CAP_DUMP
	if len(e.State) < CAP_DUMP {
		limit = len(e.State)
	}

	for i := 0; i < limit; i++ {
		flattened := e.State[i].flatten()
		flatString := ""
		for j := 0; j < len(flattened); j++ {
			flatString += fmt.Sprintf("%v ", flattened[j])
		}

		_, err := file.WriteString(flatString + "\n")
		if err != nil {
			panic(err)
		}
	}
}

func (e *Experiment) Plot(name string) {
	p := plot.New()

	text := "Plot " + name
	text += fmt.Sprintf("\nIteration: %v", len(e.State)-1)
	text += fmt.Sprintf("\nFinal State Objective Value: %v", e.State[len(e.State)-1].Value)
	text += fmt.Sprintf("\nRuntime: %v", e.GetRuntime())
	p.Title.Text = text
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Objective Function"
	p.Add(plotter.NewGrid())

	limit := CAP_PLOT
	if len(e.State) < CAP_PLOT {
		limit = len(e.State)
	}

	pts := make(plotter.XYs, limit)

	for i, cube := range e.State {
		pts[i].X = float64(i)
		pts[i].Y = float64(cube.Value)
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	fileName := "img/" + name + ".png"

	if err := p.Save(8*vg.Inch, 8*vg.Inch, fileName); err != nil {
		panic(err)
	}
}

func (e *Experiment) GetState(i int) Cube {
	return e.State[i]
}

func (e *Experiment) GetRuntime() time.Duration {
	return e.Runtime
}

func (e *Experiment) AppendState(cube *Cube) {
	e.State = append(e.State, *cube.Clone())
}

func (e *Experiment) SetRuntime(runtime time.Duration) {
	e.Runtime = runtime
}

func (e *Experiment) GetEndState() Cube {
	return e.State[len(e.State)-1]
}
