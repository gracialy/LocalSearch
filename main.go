package main

import (
	"fmt"
)

const (
	ITERATION = 3
)

func main() {
	fmt.Printf("===========================================================================================================\n")
	fmt.Printf("                                               LOCAL SEARCH\n")
	fmt.Printf("===========================================================================================================\n")
	cube := NewCube()
	cube.PrintSideways()
	fmt.Printf("===========================================================================================================\n")

	fmt.Printf("Running stochastic...\n")
	s1 := NewStochastic(cube, 1000)
	s2 := NewStochastic(cube, 2000)
	s3 := NewStochastic(cube, 5000)
	s1.Run()
	s2.Run()
	s3.Run()

	fmt.Printf("===========================================================================================================\n")
	fmt.Printf("ALGORITHM\t\tRUN\tTIME\tVALUE\tV1\tV2\tITERATION\tSTUCK\tPOPULATION\n")
	fmt.Printf("===========================================================================================================\n")
	fmt.Printf("Initial\t\t\t\t\t%d\t%d\t%d\n", cube.GetValue(), cube.GetValue1(), cube.GetValue2())

	fmt.Printf("Stochastic\t\t%d\t%.2f\t%d\t%d\t%d\t%d\n", 1, s1.GetRuntime().Seconds(), s1.GetEndState().Value, s1.GetEndState().Value1, s1.GetEndState().Value2, s1.MaxIterations)
	fmt.Printf("Stochastic\t\t%d\t%.2f\t%d\t%d\t%d\t%d\n", 2, s2.GetRuntime().Seconds(), s2.GetEndState().Value, s2.GetEndState().Value1, s2.GetEndState().Value2, s2.MaxIterations)
	fmt.Printf("Stochastic\t\t%d\t%.2f\t%d\t%d\t%d\t%d\n", 3, s3.GetRuntime().Seconds(), s3.GetEndState().Value, s3.GetEndState().Value1, s3.GetEndState().Value2, s3.MaxIterations)

	// ==========================================================================================================================================================================================================================================
	// PLOTTING
	// ==========================================================================================================================================================================================================================================

	s1.Plot("Stochastic 1")
	s2.Plot("Stochastic 2")
	s3.Plot("Stochastic 3")
}
