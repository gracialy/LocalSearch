package main

import "fmt"

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
	cubeForStochastic := cube.Copy()
	st := NewStochastic(cubeForStochastic)
	st.Run()
	st.Cube.PrintConfiguration()
	fmt.Printf("Cube Value: %d\n", st.Cube.Value)

	// cubeForSteepest := cube.Copy()
	// sa := NewSteepestAscent(cubeForSteepest)
	// sa.Run()

	// cubeForSideways := cube.Copy()
	// sm := NewSidewaysMove(cubeForSideways)
	// sm.Run()
}
