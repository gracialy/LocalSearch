package main

import (
	"fmt"
	"os"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

	for _, cube := range e.State {
		flattened := cube.flatten()
		flatString := ""
		for i := 0; i < len(flattened); i++ {
			flatString += fmt.Sprintf("%v ", flattened[i])
		}
		fmt.Println(flatString)

		_, err := file.WriteString(flatString + "\n")
		if err != nil {
			panic(err)
		}
	}
}

func (e *Experiment) Plot(name string) {
	p := plot.New()

	text := "Plot " + name
	text += fmt.Sprintf("\nFinal State Objective Value: %v", e.State[len(e.State)-1].Value)
	text += fmt.Sprintf("\nRuntime: %v", e.GetRuntime())
	p.Title.Text = text
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Objective Function"
	p.Add(plotter.NewGrid())

	pts := make(plotter.XYs, len(e.State))

	for i, cube := range e.State {
		pts[i].X = float64(i)
		pts[i].Y = float64(cube.Value)
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	if err := p.Save(8*vg.Inch, 8*vg.Inch, "img/"+text+".png"); err != nil {
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
