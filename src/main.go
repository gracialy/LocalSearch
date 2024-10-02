package main

import (
	"fmt"
)

func main() {
	fmt.Printf("========================================================================\n")
	fmt.Printf("                       LOCAL SEARCH ALGORITHM\n")
	fmt.Printf("========================================================================\n")
	cube := NewCube()
	cube.PrintSideways()

	fmt.Printf("========================================================================\n")
	fmt.Printf("                         STEEPEST ASCENT ALGORITHM\n")
	fmt.Printf("========================================================================\n")
	cubeForSTA := cube.Clone()
	sta := NewSteepestAscent(cubeForSTA)
	sta.Run()
	sta.PrintSideways()

	fmt.Printf("========================================================================\n")
	fmt.Printf("                         SIDEWAYS MOVE ALGORITHM\n")
	fmt.Printf("========================================================================\n")
	cubeForSM := cube.Clone()
	sm := NewSidewaysMove(cubeForSM)
	sm.Run()
	sm.PrintSideways()

	fmt.Printf("========================================================================\n")
	fmt.Printf("                    RANDOM RESTART (STEEPEST ASCENT) ALGORITHM\n")
	fmt.Printf("========================================================================\n")
	cubeForRR_STA := cube.Clone()
	rr_sta := NewRR_sta(cubeForRR_STA)
	rr_sta.Run()
	rr_sta.PrintSideways()

	fmt.Printf("========================================================================\n")
	fmt.Printf("                    RANDOM RESTART (SIDEWAYS MOVE) ALGORITHM\n")
	fmt.Printf("========================================================================\n")
	cubeForRR_SM := cube.Clone()
	rr_sm := NewRR_sm(cubeForRR_SM)
	rr_sm.Run()
	rr_sm.PrintSideways()

	fmt.Printf("========================================================================\n")
	fmt.Printf("                    STOCHASTIC HILL CLIMBING ALGORITHM\n")
	fmt.Printf("========================================================================\n")
	cubeForST := cube.Clone()
	st := NewStochastic(cubeForST)
	st.Run()
	st.PrintSideways()

	fmt.Printf("========================================================================\n")
	fmt.Printf("                      SIMULATED ANNEALING ALGORITHM\n")
	fmt.Printf("========================================================================\n")
	cubeForSA := cube.Clone()
	sa := NewSimulatedAnnealing(cubeForSA)
	sa.Run()
	sa.PrintSideways()

	fmt.Printf("========================================================================\n")
	fmt.Printf("ALGORITHM			RUNTIME		VALUE	V1 	V2\n")
	fmt.Printf("========================================================================\n")
	fmt.Printf("Initial				0		%d	%d	%d\n", cube.GetValue(), cube.GetValue1(), cube.GetValue2())
	fmt.Printf("Steepest Ascend 		%.2f		%d	%d	%d\n", sta.GetRuntime().Seconds(), sta.GetValue(), sta.GetValue1(), sta.GetValue2())
	fmt.Printf("Sideways Move 			%.2f		%d	%d	%d\n", sm.GetRuntime().Seconds(), sm.GetValue(), sm.GetValue1(), sm.GetValue2())
	fmt.Printf("Random Restart (StA)		%.2f		%d	%d	%d\n", rr_sta.GetRuntime().Seconds(), rr_sta.GetValue(), rr_sta.GetValue1(), rr_sta.GetValue2())
	fmt.Printf("Random Restart (SM)		%.2f		%d	%d	%d\n", rr_sm.GetRuntime().Seconds(), rr_sm.GetValue(), rr_sm.GetValue1(), rr_sm.GetValue2())
	fmt.Printf("Stochastic 			%.2f		%d	%d	%d\n", st.GetRuntime().Seconds(), st.GetValue(), st.GetValue1(), st.GetValue2())
	fmt.Printf("Simulated Annealing		%.2f		%d	%d	%d\n", sa.GetRuntime().Seconds(), sa.GetValue(), sa.GetValue1(), sa.GetValue2())
}
