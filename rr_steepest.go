package main

import (
	"fmt"
	"os"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type RR_sta struct {
	Restart       []SteepestAscent
	MaxRestart    int
	ActualRuntime time.Duration
}

func NewRR_sta(cube *Cube, maxRestart int) *RR_sta {
	rr_sta := &RR_sta{
		MaxRestart: maxRestart,
	}
	rr_sta.Restart = append(make([]SteepestAscent, 0), *NewSteepestAscent(cube))
	return rr_sta
}

func (rr_sta *RR_sta) Run() {
	start := time.Now()

	for i := 0; i < rr_sta.MaxRestart; i++ {
		if i != 0 {
			randomState := NewCube()
			sta := NewSteepestAscent(randomState)
			rr_sta.AppendRestart(sta)
		}

		rr_sta.Restart[i].Run()

		if rr_sta.Restart[i].Experiment.GetEndState().Value == 0 {
			break
		}
	}

	rr_sta.ActualRuntime = time.Since(start)
}

func (rr_sta *RR_sta) Plot(name string) {
	p := plot.New()

	text := "Plot " + name
	text += fmt.Sprintf("\nActual Restart: %v", len(rr_sta.Restart))
	text += fmt.Sprintf("\nAverage Iterations: %v", rr_sta.AverageIterations())
	text += fmt.Sprintf("\nFinal State Objective Value: %v", rr_sta.GetFinalObjectiveValue())
	text += fmt.Sprintf("\nRuntime: %v", rr_sta.GetRuntime())
	p.Title.Text = text
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Objective Function"
	p.Add(plotter.NewGrid())

	var pts plotter.XYs

	count := 0
	for i, restart := range rr_sta.Restart {
		for j, state := range restart.Experiment.State {
			if count > CAP_PLOT {
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

func (rr_sta *RR_sta) IterationPlot(name string) {
	p := plot.New()

	text := "Iteration Plot " + name
	p.Title.Text = text
	p.X.Label.Text = "Restart"
	p.Y.Label.Text = "Iteration"
	p.Add(plotter.NewGrid())

	limit := CAP_PLOT
	if len(rr_sta.Restart) < CAP_PLOT {
		limit = len(rr_sta.Restart)
	}

	pts := make(plotter.XYs, limit)

	for i := 0; i < limit; i++ {
		pts[i].X = float64(i)
		pts[i].Y = float64(rr_sta.Restart[i].ActualIteration)
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	// scatter, err := plotter.NewScatter(pts)
	// if err != nil {
	// 	panic(err)
	// }
	// scatter.GlyphStyle.Shape = draw.CircleGlyph{}
	// p.Add(scatter)

	if err := p.Save(8*vg.Inch, 8*vg.Inch, "img/"+text+".png"); err != nil {
		panic(err)
	}
}

func (rr_sta *RR_sta) Dump(name string) {
	file, err := os.Create("txt/" + name + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	count := 0
	for _, restart := range rr_sta.Restart {
		for _, state := range restart.Experiment.State {
			if count > CAP_DUMP {
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

func (rr_sta *RR_sta) GetRuntime() time.Duration {
	return rr_sta.ActualRuntime
}

func (rr_sta *RR_sta) GetEndState() *Cube {
	return rr_sta.Restart[len(rr_sta.Restart)-1].Experiment.GetEndState()
}

func (rr_sta *RR_sta) GetFinalObjectiveValue() int {
	return rr_sta.Restart[len(rr_sta.Restart)-1].Experiment.GetEndState().Value
}

func (rr_sta *RR_sta) AppendRestart(sta *SteepestAscent) {
	rr_sta.Restart = append(rr_sta.Restart, *sta)
}

func (rr_sta *RR_sta) AverageIterations() int {
	total := 0
	for _, sta := range rr_sta.Restart {
		total += sta.ActualIteration
	}
	return total / len(rr_sta.Restart)
}
