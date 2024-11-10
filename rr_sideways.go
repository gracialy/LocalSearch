package main

import (
	"fmt"
	"os"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type RR_sm struct {
	Restart       []SidewaysMove
	MaxRestart    int
	MaxSideways   int
	ActualRuntime time.Duration
}

func NewRR_sm(cube *Cube, maxRestart int, maxSideways int) *RR_sm {
	rr_sm := &RR_sm{
		MaxRestart:  maxRestart,
		MaxSideways: maxSideways,
	}
	rr_sm.Restart = append(make([]SidewaysMove, 0), *NewSidewaysMove(cube, maxSideways))
	return rr_sm
}

func (rr_sm *RR_sm) Run() {
	start := time.Now()

	for i := 0; i < rr_sm.MaxRestart; i++ {
		if i != 0 {
			randomState := NewCube()
			sm := NewSidewaysMove(randomState, rr_sm.MaxSideways)
			rr_sm.AppendRestart(sm)
		}

		rr_sm.Restart[i].Run()

		if rr_sm.Restart[i].Experiment.GetEndState().Value == 0 {
			break
		}
	}

	rr_sm.ActualRuntime = time.Since(start)
}

func (rr_sm *RR_sm) Plot(name string) {
	p := plot.New()

	text := "Plot " + name
	text += fmt.Sprintf("\nActual Restart: %v", len(rr_sm.Restart))
	text += fmt.Sprintf("\nAverage Iterations: %v", rr_sm.AverageIterations())
	text += fmt.Sprintf("\nFinal State Objective Value: %v", rr_sm.GetFinalObjectiveValue())
	text += fmt.Sprintf("\nRuntime: %v", rr_sm.GetRuntime())
	text += fmt.Sprintf("\nMax Sideways: %v", rr_sm.MaxSideways)
	p.Title.Text = text
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Objective Function"
	p.Add(plotter.NewGrid())

	var pts plotter.XYs

	count := 0
	for i, restart := range rr_sm.Restart {
		for j, state := range restart.Experiment.State {
			if count >= CAP_PLOT {
				break
			}
			count++
			pts = append(pts, plotter.XY{X: float64(i*len(restart.Experiment.State) + j), Y: float64(state.Value)})
		}
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

func (rr_sm *RR_sm) IterationPlot(name string) {
	p := plot.New()

	text := "Iteration Plot " + name
	p.Title.Text = text
	p.X.Label.Text = "Restart"
	p.Y.Label.Text = "Iteration"
	p.Add(plotter.NewGrid())

	limit := CAP_PLOT
	if len(rr_sm.Restart) < CAP_PLOT {
		limit = len(rr_sm.Restart)
	}

	pts := make(plotter.XYs, limit)

	for i := 0; i < limit; i++ {
		pts[i].X = float64(i)
		pts[i].Y = float64(rr_sm.Restart[i].ActualIteration)
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

func (rr_sm *RR_sm) Dump(name string) {
	file, err := os.Create("txt/" + name + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	count := 0
	for _, restart := range rr_sm.Restart {
		for _, state := range restart.Experiment.State {
			if count >= CAP_DUMP {
				break
			}
			count++
			flattened := state.flatten()
			flatString := ""
			for _, value := range flattened {
				flatString += fmt.Sprintf("%v ", value)
			}

			_, err := file.WriteString(flatString + "\n")
			if err != nil {
				panic(err)
			}
		}
	}
}

func (rr_sm *RR_sm) GetRuntime() time.Duration {
	return rr_sm.ActualRuntime
}

func (rr_sm *RR_sm) GetEndState() Cube {
	return rr_sm.Restart[len(rr_sm.Restart)-1].Experiment.GetEndState()
}

func (rr_sm *RR_sm) GetFinalObjectiveValue() int {
	return rr_sm.Restart[len(rr_sm.Restart)-1].Experiment.GetEndState().Value
}

func (rr_sm *RR_sm) AppendRestart(sm *SidewaysMove) {
	rr_sm.Restart = append(rr_sm.Restart, *sm)
}

func (rr_sm *RR_sm) AverageIterations() int {
	total := 0
	for _, sm := range rr_sm.Restart {
		total += sm.ActualIteration
	}
	return total / len(rr_sm.Restart)
}
