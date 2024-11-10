package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const (
	BIG_ELITE      = 0.05
	SMALL_ELITE    = 0.1
	ELITE_SPAN     = 1
	BIG_MUTATION   = 0.05
	SMALL_MUTATION = 0.01
)

var eliteDeath int
var eliteAge int = 0
var elitePercent float64
var eliteSize int
var mutationProb float64

type GeneticAlgorithm struct {
	Experiment
	PopulationSize  int
	Population      []Cube
	MaxIterations   int
	ActualIteration int
	AvgObjective    []int
}

func NewGeneticAlgorithm(cube *Cube, populationSize int, maxIterations int) *GeneticAlgorithm {
	ga := &GeneticAlgorithm{}

	ga.PopulationSize = populationSize
	ga.MaxIterations = maxIterations
	ga.ActualIteration = 0
	ga.AvgObjective = []int{}

	ga.Population = make([]Cube, ga.PopulationSize)

	ga.Population[0] = *cube.Clone()
	for i := 1; i < ga.PopulationSize; i++ {
		ga.Population[i] = *NewCube()
	}

	ga.Experiment = *NewExperiment(cube.Clone())

	elitePercent = BIG_ELITE
	if populationSize < 200 {
		elitePercent = SMALL_ELITE
	}

	eliteDeath = int(ELITE_SPAN * float64(maxIterations))
	eliteSize = int(float64(populationSize) * elitePercent)

	return ga
}

func (ga *GeneticAlgorithm) Run() {
	start := time.Now()
	eliteAge = 0
	eliteSize = int(float64(ga.PopulationSize) * elitePercent)

	// clear the dump file
	// err := os.WriteFile("tmp/ga_dump.txt", []byte{}, 0644)
	// if err != nil {
	// 	fmt.Println("Error clearing tmp/ga_dump.txt:", err)
	// 	return
	// }
	// fmt.Printf("\n")
	//

	for {
		ga.Sort()

		// keep track of the population value
		// populationValues := ""
		// for i := 0; i < ga.PopulationSize; i++ {
		// 	populationValues += fmt.Sprintf("%d ", ga.Population[i].Value)
		// }
		// populationValues += "\n"

		// f, err := os.OpenFile("tmp/ga_dump.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		// if err != nil {
		// 	fmt.Println("Error opening tmp/ga_dump.txt:", err)
		// 	return
		// }
		// defer f.Close()

		// if _, err = f.WriteString(populationValues); err != nil {
		// 	fmt.Println("Error writing to tmp/ga_dump.txt:", err)
		// 	return
		// }
		//

		// add best cube to the experiment state
		if ga.ActualIteration == 0 {
			ga.Experiment.State[0] = ga.Population[0]
		} else {
			ga.Experiment.AppendState(&ga.Population[0])
		}

		// calculate average objective value
		avg := 0
		for i := 0; i < ga.PopulationSize; i++ {
			avg += ga.Population[i].Value
		}
		avg /= ga.PopulationSize
		ga.AvgObjective = append(ga.AvgObjective, avg)

		// break condition
		if ga.Population[0].Value == 0 || ga.EndSearch() {
			break
		}

		ga.NextGeneration()
	}

	ga.Experiment.SetRuntime(time.Since(start))
}

func (ga *GeneticAlgorithm) Sort() {
	for i := 0; i < ga.PopulationSize; i++ {
		for j := i + 1; j < ga.PopulationSize; j++ {
			if ga.Population[i].Value > ga.Population[j].Value {
				ga.Population[i], ga.Population[j] = ga.Population[j], ga.Population[i]
			}
		}
	}
}

func (ga *GeneticAlgorithm) NextGeneration() {
	ga.ActualIteration++
	nextPopulation := make([]Cube, 0)

	if eliteAge < eliteDeath {
		eliteAge++
		nextPopulation = append(nextPopulation, ga.Population[:eliteSize]...)
	} else {
		eliteAge = 0
	}

	// selection wheel
	selectionWheel := ga.CreateSelectionWheel()
	// fmt.Println("Selection Wheel:", selectionWheel)

	for {
		if len(nextPopulation) == ga.PopulationSize {
			break
		}

		// selection
		parent1, parent2 := ga.SelectParents(selectionWheel)
		// crossover
		child1, child2 := ga.Crossover(parent1, parent2)
		// mutation
		ga.Mutate(child1)
		ga.Mutate(child2)

		if !ga.IsDuplicate(*child1) && !ga.IsDuplicate(*child2) {
			nextPopulation = append(nextPopulation, *child1)
		}

		if !ga.IsDuplicate(*child2) && len(nextPopulation) < ga.PopulationSize {
			nextPopulation = append(nextPopulation, *child2)
		}
	}

	ga.Population = nextPopulation
}

func (ga *GeneticAlgorithm) CreateSelectionWheel() []float64 {
	selectionWheel := make([]float64, 0)

	sum := 0
	for i := 0; i < len(ga.Population); i++ {
		sum += ga.Population[i].Value
	}
	position := 0.0
	for i := 0; i < len(ga.Population); i++ {
		position += float64(ga.Population[i].Value) / float64(sum)
		selectionWheel = append(selectionWheel, position)
	}

	return selectionWheel
}

func (ga *GeneticAlgorithm) SelectParents(selectionWheel []float64) (Cube, Cube) {
	parent1 := ga.Population[0]
	parent2 := ga.Population[0]

	prob1 := rand.Float64()
	for i := 0; i < len(selectionWheel); i++ {
		if prob1 <= selectionWheel[i] {
			parent1 = ga.Population[i]
			break
		}
	}

	for {
		prob2 := rand.Float64()
		for i := 0; i < len(selectionWheel); i++ {
			if prob2 <= selectionWheel[i] {
				parent2 = ga.Population[i]
				break
			}
		}
		if parent1 != parent2 {
			break
		}
	}

	return parent1, parent2
}

func (ga *GeneticAlgorithm) Crossover(parent1 Cube, parent2 Cube) (*Cube, *Cube) {
	parent1 = *parent1.Clone()
	parent2 = *parent2.Clone()
	parentFlat1 := parent1.flatten()
	parentFlat2 := parent2.flatten()

	child1 := NewCube()
	child2 := NewCube()

	// order crossover (ox1)
	var point1, point2 int
	for {
		point1 = rand.Intn(ELEMENT)
		point2 = rand.Intn(ELEMENT)
		if point1 != point2 {
			if point1 > point2 {
				point1, point2 = point2, point1
			}
			break
		}
	}

	childFlat1 := make([]uint8, ELEMENT)
	childFlat2 := make([]uint8, ELEMENT)

	// initialize child with 0 to indicate empty slots
	for i := 0; i < ELEMENT; i++ {
		childFlat1[i] = uint8(0)
		childFlat2[i] = uint8(0)
	}

	// copy the genes within points range from respective parents
	for i := point1; i <= point2; i++ {
		childFlat1[i] = parentFlat1[i]
		childFlat2[i] = parentFlat2[i]
	}

	// get the list of the other parent's genes that are not in the points range
	remainingGenes1 := []uint8{}
	remainingGenes2 := []uint8{}
	for i := 0; i < ELEMENT; i++ {
		if !contains(childFlat1, parentFlat2[i]) {
			remainingGenes1 = append(remainingGenes1, parentFlat2[i])
		}
		if !contains(childFlat2, parentFlat1[i]) {
			remainingGenes2 = append(remainingGenes2, parentFlat1[i])
		}
	}

	// fill the remaining genes in child
	for i := 0; i < ELEMENT; i++ {
		if childFlat1[i] == 0 {
			childFlat1[i] = remainingGenes1[0]
			remainingGenes1 = remainingGenes1[1:]
		}
		if childFlat2[i] == 0 {
			childFlat2[i] = remainingGenes2[0]
			remainingGenes2 = remainingGenes2[1:]
		}
	}

	child1.unflatten(childFlat1)
	child2.unflatten(childFlat2)

	return child1, child2
}

func contains(arr []uint8, val uint8) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func (ga *GeneticAlgorithm) Mutate(cube *Cube) {
	if ga.IsGettingConvergen() {
		mutationProb = BIG_MUTATION
	} else {
		mutationProb = SMALL_MUTATION
	}

	if rand.Float64() >= mutationProb {
		return
	}

	cubeFlat := cube.flatten()

	// inversion
	i := rand.Intn(len(cubeFlat))
	j := i
	for i == j {
		d := rand.Float64() * float64(len(cubeFlat)-1)
		j = (i + int(d)) % len(cubeFlat)
	}

	// keep track of the selected genes to invert
	selectedGenes := make([]uint8, 0)
	if i <= j {
		selectedGenes = append(selectedGenes, cubeFlat[i:j+1]...)
	} else {
		selectedGenes = append(selectedGenes, cubeFlat[i:]...)
		selectedGenes = append(selectedGenes, cubeFlat[:j+1]...)
	}
	// fmt.Printf("Selected genes to invert: %v\n", selectedGenes)

	// invert the selected genes
	for k := 0; k < len(selectedGenes)/2; k++ {
		selectedGenes[k], selectedGenes[len(selectedGenes)-1-k] = selectedGenes[len(selectedGenes)-1-k], selectedGenes[k]
	}

	if i < j {
		for k := i; k <= j; k++ {
			cubeFlat[k] = selectedGenes[k-i]
		}
	} else {
		for k := i; k < len(cubeFlat); k++ {
			cubeFlat[k] = selectedGenes[k-i]
		}
		for k := 0; k <= j; k++ {
			cubeFlat[k] = selectedGenes[len(cubeFlat)-i+k]
		}
	}

	cube.unflatten(cubeFlat)
}

func (ga *GeneticAlgorithm) IsGettingConvergen() bool {
	if ga.ActualIteration < 50 {
		return false
	}

	for i := 1; i < 50; i++ {
		if ga.AvgObjective[ga.ActualIteration-i] != ga.AvgObjective[ga.ActualIteration-1] {
			return false
		}
	}

	return true
}

func (ga *GeneticAlgorithm) EndSearch() bool {
	return ga.ActualIteration >= ga.MaxIterations
}

func (ga *GeneticAlgorithm) IsDuplicate(cube Cube) bool {
	for i := 0; i < len(ga.Population); i++ {
		if cube.IsSame(&ga.Population[i]) {
			return true
		}
	}
	return false
}

func (ga *GeneticAlgorithm) Plot(name string) {
	p := plot.New()
	e := &ga.Experiment

	text := "Plot " + name
	text += fmt.Sprintf("\nIteration: %v", ga.ActualIteration)
	text += fmt.Sprintf("\nFinal State Objective Value: %v", e.State[len(e.State)-1].Value)
	text += fmt.Sprintf("\nRuntime: %v", e.GetRuntime())
	text += fmt.Sprintf("\nPopulation Size: %v", ga.PopulationSize)
	text += fmt.Sprintf("\nMax Iterations: %v", ga.MaxIterations)
	p.Title.Text = text
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Objective Function"
	p.Add(plotter.NewGrid())

	limit := CAP_DUMP
	if len(e.State) < CAP_DUMP {
		limit = len(e.State)
	}

	pts := make(plotter.XYs, limit)
	avgPts := make(plotter.XYs, limit)

	for i := 0; i < limit; i++ {
		pts[i].X = float64(i)
		pts[i].Y = float64(e.State[i].Value)
		avgPts[i].X = float64(i)
		avgPts[i].Y = float64(ga.AvgObjective[i])
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	avgLine, err := plotter.NewLine(avgPts)
	if err != nil {
		panic(err)
	}
	avgLine.LineStyle.Color = plotutil.Color(0)
	p.Add(avgLine)

	p.Legend.Add("Best", line)
	p.Legend.Add("Average", avgLine)

	fileName := "img/" + name + ".png"

	if err := p.Save(8*vg.Inch, 8*vg.Inch, fileName); err != nil {
		panic(err)
	}
}
