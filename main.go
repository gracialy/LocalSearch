package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cube *Cube
var secondGA int = 0

// var triggered bool

func main() {
	runExperiment()
}

func runExperiment() {
	for {
		cube = NewCube()
		// triggered = false
		fmt.Printf("===========================================================================================================\n")
		fmt.Printf("                                               LOCAL SEARCH\n")
		fmt.Printf("===========================================================================================================\n")
		fmt.Println("Choose an experiment to perform:")
		fmt.Println("1. Steepest Ascent")
		fmt.Println("2. Sideways Move")
		fmt.Println("3. Random Restart (Steepest Ascent)")
		fmt.Println("4. Random Restart (Sideways Move)")
		fmt.Println("5. Stochastic")
		fmt.Println("6. Simulated Annealing")
		fmt.Println("7. Genetic Algorithm")
		fmt.Println("8. Benchmark All")
		fmt.Println("9. Exit")

		choice := atoi(getUserInput("Enter your choice: "))

		switch choice {
		case 1:
			runSteepestAscent()
		case 2:
			runSidewaysMove()
		case 3:
			runRandomRestartSteepestAscent()
		case 4:
			runRandomRestartSidewaysMove()
		case 5:
			runStochastic()
		case 6:
			runSimulatedAnnealing()
		case 7:
			runGeneticAlgorithm()
		case 8:
			benchmarkAll()
		case 9:
			fmt.Printf("Thank you for using this program.\n")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func runSteepestAscent() {
	fmt.Printf("\n===STEEPEST ASCENT===\n")
	fmt.Printf("Running steepest ascent...\n")
	sta1 := NewSteepestAscent(cube)
	sta2 := NewSteepestAscent(cube)
	sta3 := NewSteepestAscent(cube)
	sta1.Run()
	sta2.Run()
	sta3.Run()

	resultHeader()
	fmt.Printf("Steepest Ascent\t\t%d\t%.2f\t%d\t%d\t%d\t%d\n", 1, sta1.GetRuntime().Seconds(), sta1.GetEndState().Value, sta1.GetEndState().Value1, sta1.GetEndState().Value2, sta1.ActualIteration)
	fmt.Printf("Steepest Ascent\t\t%d\t%.2f\t%d\t%d\t%d\t%d\n", 2, sta2.GetRuntime().Seconds(), sta2.GetEndState().Value, sta2.GetEndState().Value1, sta2.GetEndState().Value2, sta2.ActualIteration)
	fmt.Printf("Steepest Ascent\t\t%d\t%.2f\t%d\t%d\t%d\t%d\n", 3, sta3.GetRuntime().Seconds(), sta3.GetEndState().Value, sta3.GetEndState().Value1, sta3.GetEndState().Value2, sta3.ActualIteration)

	fmt.Printf("Generating dump file...\n")
	sta1.Dump("Steepest Ascent 1")
	sta2.Dump("Steepest Ascent 2")
	sta3.Dump("Steepest Ascent 3")

	fmt.Printf("Generating plot file...\n")
	sta1.Plot("Steepest Ascent 1")
	sta2.Plot("Steepest Ascent 2")
	sta3.Plot("Steepest Ascent 3")
}

func runSidewaysMove() {
	fmt.Printf("\n===SIDEWAYS MOVE===\n")
	maxSideways := []int{
		atoi(getUserInput("Enter max sideways moves for sm1: ")),
		atoi(getUserInput("Enter max sideways moves for sm2: ")),
		atoi(getUserInput("Enter max sideways moves for sm3: ")),
	}

	fmt.Printf("Running sideways move...\n")
	sm1 := NewSidewaysMove(cube, maxSideways[0])
	sm2 := NewSidewaysMove(cube, maxSideways[1])
	sm3 := NewSidewaysMove(cube, maxSideways[2])
	sm1.Run()
	sm2.Run()
	sm3.Run()

	resultHeader()
	fmt.Printf("Sideways Move\t\t%d\t%.2f\t%d\t%d\t%d\t%d\n", 1, sm1.GetRuntime().Seconds(), sm1.GetEndState().Value, sm1.GetEndState().Value1, sm1.GetEndState().Value2, sm1.ActualIteration)
	fmt.Printf("Sideways Move\t\t%d\t%.2f\t%d\t%d\t%d\t%d\n", 2, sm2.GetRuntime().Seconds(), sm2.GetEndState().Value, sm2.GetEndState().Value1, sm2.GetEndState().Value2, sm2.ActualIteration)
	fmt.Printf("Sideways Move\t\t%d\t%.2f\t%d\t%d\t%d\t%d\n", 3, sm3.GetRuntime().Seconds(), sm3.GetEndState().Value, sm3.GetEndState().Value1, sm3.GetEndState().Value2, sm3.ActualIteration)

	fmt.Printf("Generating dump file...\n")
	sm1.Dump("Sideways Move 1")
	sm2.Dump("Sideways Move 2")
	sm3.Dump("Sideways Move 3")

	fmt.Printf("Generating plot file...\n")
	sm1.Plot("Sideways Move 1")
	sm2.Plot("Sideways Move 2")
	sm3.Plot("Sideways Move 3")
}

func runRandomRestartSteepestAscent() {
	fmt.Printf("\n===RANDOM RESTART (STEEPEST ASCENT)===\n")
}

func runRandomRestartSidewaysMove() {
	fmt.Printf("\n===RANDOM RESTART (SIDEWAYS MOVE)===\n")
}

func runStochastic() {
	fmt.Printf("\n===STOCHASTIC===\n")
	maxIterations := []int{
		atoi(getUserInput("Enter max iterations for s1: ")),
		atoi(getUserInput("Enter max iterations for s2: ")),
		atoi(getUserInput("Enter max iterations for s3: ")),
	}

	fmt.Println("Running stochastic...")
	s1 := NewStochastic(cube, maxIterations[0])
	s2 := NewStochastic(cube, maxIterations[1])
	s3 := NewStochastic(cube, maxIterations[2])
	s1.Run()
	s2.Run()
	s3.Run()

	resultHeader()
	fmt.Printf("Stochastic\t\t%d\t%.2f\t%d\t%d\t%d\t%d\n", 1, s1.GetRuntime().Seconds(), s1.GetEndState().Value, s1.GetEndState().Value1, s1.GetEndState().Value2, s1.MaxIterations)
	fmt.Printf("Stochastic\t\t%d\t%.2f\t%d\t%d\t%d\t%d\n", 2, s2.GetRuntime().Seconds(), s2.GetEndState().Value, s2.GetEndState().Value1, s2.GetEndState().Value2, s2.MaxIterations)
	fmt.Printf("Stochastic\t\t%d\t%.2f\t%d\t%d\t%d\t%d\n", 3, s3.GetRuntime().Seconds(), s3.GetEndState().Value, s3.GetEndState().Value1, s3.GetEndState().Value2, s3.MaxIterations)

	fmt.Printf("Generating dump file...\n")
	s1.Dump("Stochastic 1")
	s2.Dump("Stochastic 2")
	s3.Dump("Stochastic 3")

	fmt.Printf("Generating plot file...\n")
	s1.Plot("Stochastic 1")
	s2.Plot("Stochastic 2")
	s3.Plot("Stochastic 3")
}

func runSimulatedAnnealing() {
	fmt.Printf("\n===SIMULATED ANNEALING===\n")
}

func runGeneticAlgorithm() {
	fmt.Printf("===GENETIC ALGORITHM===\n")
	multiplier := 1 + secondGA
	populationSize := []int{
		atoi(getUserInput(fmt.Sprintf("Enter POPULATION SIZE for ga%v: ", 1*multiplier))),
		atoi(getUserInput(fmt.Sprintf("Enter POPULATION SIZE for ga%v: ", 2*multiplier))),
		atoi(getUserInput(fmt.Sprintf("Enter POPULATION SIZE for ga%v: ", 3*multiplier))),
	}
	maxIterations := []int{
		atoi(getUserInput(fmt.Sprintf("Enter MAX ITERATIONS for ga%v: ", 1*multiplier))),
		atoi(getUserInput(fmt.Sprintf("Enter MAX ITERATIONS for ga%v: ", 2*multiplier))),
		atoi(getUserInput(fmt.Sprintf("Enter MAX ITERATIONS for ga%v: ", 3*multiplier))),
	}

	fmt.Printf("Running genetic algorithm...\n")
	ga1 := NewGeneticAlgorithm(cube, populationSize[0], maxIterations[0])
	ga2 := NewGeneticAlgorithm(cube, populationSize[1], maxIterations[1])
	ga3 := NewGeneticAlgorithm(cube, populationSize[2], maxIterations[2])
	ga1.Run()
	ga2.Run()
	ga3.Run()

	fmt.Printf("Generating dump file...\n")
	ga1.Dump("Genetic Algorithm" + strconv.Itoa(1*multiplier))
	ga2.Dump("Genetic Algorithm" + strconv.Itoa(2*multiplier))
	ga3.Dump("Genetic Algorithm" + strconv.Itoa(3*multiplier))

	fmt.Printf("Generating plot file...\n")
	ga1.Plot("Genetic Algorithm" + strconv.Itoa(1*multiplier))
	ga2.Plot("Genetic Algorithm" + strconv.Itoa(2*multiplier))
	ga3.Plot("Genetic Algorithm" + strconv.Itoa(3*multiplier))

	resultHeader()
	resultHeader()
	fmt.Printf("Genetic Algorithm\t%d\t%.2f\t%d\t%d\t%d\t%d\t\t\t%d\n", 1*multiplier, ga1.GetRuntime().Seconds(), ga1.GetEndState().Value, ga1.GetEndState().Value1, ga1.GetEndState().Value2, ga1.ActualIteration, ga1.PopulationSize)
	fmt.Printf("Genetic Algorithm\t%d\t%.2f\t%d\t%d\t%d\t%d\t\t\t%d\n", 2*multiplier, ga2.GetRuntime().Seconds(), ga2.GetEndState().Value, ga2.GetEndState().Value1, ga2.GetEndState().Value2, ga2.ActualIteration, ga2.PopulationSize)
	fmt.Printf("Genetic Algorithm\t%d\t%.2f\t%d\t%d\t%d\t%d\t\t\t%d\n", 3*multiplier, ga3.GetRuntime().Seconds(), ga3.GetEndState().Value, ga3.GetEndState().Value1, ga3.GetEndState().Value2, ga3.ActualIteration, ga3.PopulationSize)
}

func benchmarkAll() {
	runSteepestAscent()
	runSidewaysMove()
	runRandomRestartSteepestAscent()
	runRandomRestartSidewaysMove()
	runStochastic()
	runSimulatedAnnealing()
	runSimulatedAnnealing()
	secondGA = 1
	runGeneticAlgorithm()
	secondGA = 0
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func atoi(input string) int {
	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Invalid input: %s\n", input)
		os.Exit(1)
	}
	return value
}

func atof(input string) float64 {
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("Invalid input: %s\n", input)
		os.Exit(1)
	}
	return value
}

func resultHeader() {
	// if triggered {
	// 	return
	// }
	fmt.Printf("===========================================================================================================\n")
	fmt.Printf("ALGORITHM\t\tRUN\tTIME\tVALUE\tV1\tV2\tITERATION\tSTUCK\tPOPULATION\n")
	fmt.Printf("===========================================================================================================\n")
	fmt.Printf("Initial\t\t\t\t\t%d\t%d\t%d\n", cube.GetValue(), cube.GetValue1(), cube.GetValue2())
	// triggered = true
}
