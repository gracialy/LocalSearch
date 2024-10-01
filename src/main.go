package main

import (
	"fmt"
)

func main() {
	fmt.Printf("====================================================\n")
	fmt.Printf("               LOCAL SEARCH ALGORITHM\n")
	fmt.Printf("====================================================\n")
	cube := NewCube()
	cube.PrintConfiguration()
	fmt.Printf("Cube Value: %d\n\n", cube.Value)

	fmt.Printf("====================================================\n")
	fmt.Printf("         STOCHASTIC HILL CLIMBING ALGORITHM\n")
	fmt.Printf("====================================================\n")
	cubeForST := cube.Clone()
	st := NewStochastic(cubeForST)
	st.Run()
	st.PrintConfiguration()
	fmt.Printf("Cube Value: %d\n\n", st.GetValue())

	fmt.Printf("====================================================\n")
	fmt.Printf("          SIMULATED ANNEALING ALGORITHM\n")
	fmt.Printf("====================================================\n")
	cubeForSA := cube.Clone()
	sa := NewSimulatedAnnealing(cubeForSA)
	sa.Run()
	sa.PrintConfiguration()
	fmt.Printf("Cube Value: %d\n\n", sa.GetValue())

	fmt.Printf("====================================================\n")
	fmt.Printf("ALGORITHM			VALUE	RUNTIME		ACTUAL ITERATION\n")
	fmt.Printf("Stochastic 			%d	%.2f\n", st.GetValue(), st.GetRuntime().Seconds())
	fmt.Printf("Simulated Annealing		%d	%.2f		%d\n", sa.GetValue(), sa.GetRuntime().Seconds(), sa.GetActualIteration())

	// cubeForSteepest := cube.Clone()
	// sa := NewSteepestAscent(cubeForSteepest)
	// sa.Run()

	// cubeForSideways := cube.Clone()
	// sm := NewSidewaysMove(cubeForSideways)
	// sm.Run()
}
