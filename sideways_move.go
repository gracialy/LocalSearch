package main

import (
	"fmt"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type SidewaysMove struct {
	Experiment
	ActualIteration int
	MaxSideways     int
}

func NewSidewaysMove(cube *Cube, maxSideways int) *SidewaysMove {
	sm := &SidewaysMove{
		MaxSideways: maxSideways,
	}
	sm.Experiment = *NewExperiment(cube.Clone())
	sm.ActualIteration = 0
	return sm
}

func (sm *SidewaysMove) Run() {
	start := time.Now()
	init := sm.Experiment.GetState(0)
	current := init.Clone()
	i := 0
	sideways := 0

	for {

		neighbor := current.FindBestNeighbor()
		if neighbor.Value > current.Value {
			break
		} else if neighbor.Value == current.Value {
			sideways++
			if sideways > sm.MaxSideways {
				break
			}
		} else {
			sideways = 0
		}
		current.Copy(neighbor)

		i++
		sm.ActualIteration = i
		sm.Experiment.AppendState(current)
	}

	sm.Experiment.SetRuntime(time.Since(start))
}

func (sm *SidewaysMove) Plot(name string) {
	p := plot.New()
	e := sm.Experiment

	text := "Plot " + name
	text += fmt.Sprintf("\nFinal State Objective Value: %v", e.State[len(e.State)-1].Value)
	text += fmt.Sprintf("\nRuntime: %v", e.GetRuntime())
	text += fmt.Sprintf("\nMax Sideways: %v", sm.MaxSideways)
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

	fileName := "img/" + name + ".png"

	if err := p.Save(8*vg.Inch, 8*vg.Inch, fileName); err != nil {
		panic(err)
	}
}
