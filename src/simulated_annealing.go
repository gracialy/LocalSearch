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

	// clearing Boltzmann track
	// err := os.WriteFile("tmp/sa_dump.txt", []byte{}, 0644)
	// if err != nil {
	// 	fmt.Println("Error clearing tmp/sa_dump.txt:", err)
	// 	return
	// }

	for i := 1; i < SA_MAX; i++ {
		sa.schedule(i)
		sa.ActualIteration = i

		if sa.T <= CAP_T {
			// very close to 0 since the T will never touch sharp 0
			break
		}

		neighbor.Copy(current)
		neighbor.FindRandomNeighbor()
		delta := neighbor.Value - current.Value
		probability := sa.probability(float64(delta))
		random := rand.Float64()
		sa.AppendProbability(probability)

		if delta < 0 {
			current.Copy(neighbor)
			sa.Boltzmann[len(sa.Boltzmann)-1] = sa.InitialT
		} else if probability > random {
			current.Copy(neighbor)
			sa.stuck++
		}

		sa.Experiment.AppendState(current)

		// keep track of the Boltzmann
		// f, err := os.OpenFile("tmp/sa_dump.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		// if err != nil {
		// 	fmt.Println("Error opening tmp/sa_dump.txt:", err)
		// 	return
		// }
		// defer f.Close()

		// if _, err := f.WriteString(fmt.Sprintf("(%6d) T: %.2f B: %.2f random: %.2f d: %3d obj: %d\n", i, sa.T, probability, random, delta, current.Value)); err != nil {
		// 	fmt.Println("Error writing tmp/sa_dump.txt:", err)
		// 	return
		// }
	}

	sa.Experiment.SetRuntime(time.Since(start))
}

func (sa *SimulatedAnnealing) schedule(t int) {
	sa.T = sa.InitialT * math.Pow(COOLING_RATE, float64(t))
}

func (sa *SimulatedAnnealing) probability(delta float64) float64 {
	return math.Exp(-delta / sa.T)
}

func (sa *SimulatedAnnealing) Plot(name string) {
	p := plot.New()
	e := sa.Experiment

	text := "Plot " + name
	text += fmt.Sprintf("\nIteration: %v", len(e.State))
	text += fmt.Sprintf("\nFinal State Objective Value: %v", e.State[len(e.State)-1].Value)
	text += fmt.Sprintf("\nRuntime: %v", e.GetRuntime())
	text += fmt.Sprintf("\nStuck: %v", sa.stuck)

	p.Title.Text = text
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Objective Function"
	p.Add(plotter.NewGrid())

	limit := CAP_PLOT
	if len(e.State) < CAP_PLOT {
		limit = len(e.State)
	}

	pts := make(plotter.XYs, limit)

	for i := 0; i < limit; i++ {
		pts[i].X = float64(i)
		pts[i].Y = float64(sa.State[i].Value) // TODO WAS WRONG
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	fileName := "../img/" + name + ".png"

	if err := p.Save(8*vg.Inch, 8*vg.Inch, fileName); err != nil {
		panic(err)
	}
}

func (sa *SimulatedAnnealing) BoltzmannPlot(name string) {
	p := plot.New()

	text := "Boltzmann Plot " + name
	p.Title.Text = text
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Boltzmann Function"
	p.Add(plotter.NewGrid())

	limit := CAP_BOLTZMANN_PLOT
	if len(sa.Boltzmann) < CAP_BOLTZMANN_PLOT {
		limit = len(sa.Boltzmann)
	}

	pts := make(plotter.XYs, limit)

	for i := 0; i < limit; i++ {
		pts[i].X = float64(i)
		pts[i].Y = float64(sa.Boltzmann[i])
	}

	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}
	p.Add(scatter)

	if err := p.Save(8*vg.Inch, 8*vg.Inch, "../img/"+text+".png"); err != nil {
		panic(err)
	}
}

func (sa *SimulatedAnnealing) AppendProbability(probability float64) {
	sa.Boltzmann = append(sa.Boltzmann, probability)
}
