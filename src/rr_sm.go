package main

import (
	"time"
)

type RR_sm struct {
	SidewaysMove
}

func NewRR_sm(cube *Cube) *RR_sm {
	return &RR_sm{
		SidewaysMove: *NewSidewaysMove(cube),
	}
}

func (rr_sm *RR_sm) GetConfiguration() [DIMENSION][DIMENSION][DIMENSION]uint8 {
	return rr_sm.SidewaysMove.GetConfiguration()
}

func (rr_sm *RR_sm) GetValue() int {
	return rr_sm.SidewaysMove.GetValue()
}

func (rr_sm *RR_sm) GetValue1() int {
	return rr_sm.SidewaysMove.GetValue1()
}

func (rr_sm *RR_sm) GetValue2() int {
	return rr_sm.SidewaysMove.GetValue2()
}

func (rr_sm *RR_sm) GetCube() *Cube {
	return rr_sm.SidewaysMove.GetCube()
}

func (rr_sm *RR_sm) GetRuntime() time.Duration {
	return rr_sm.SidewaysMove.GetRuntime()
}

func (rr_sm *RR_sm) SetValue() {
	rr_sm.SidewaysMove.SetValue()
}

func (rr_sm *RR_sm) SetRuntime(runtime time.Duration) {
	rr_sm.SidewaysMove.SetRuntime(runtime)
}

func (rr_sm *RR_sm) Clone() *RR_sm {
	return &RR_sm{SidewaysMove: *rr_sm.SidewaysMove.Clone()}
}

func (rr_sm *RR_sm) Copy(original *RR_sm) {
	rr_sm.SidewaysMove.Copy(&original.SidewaysMove)
}

func (rr_sm *RR_sm) PrintSideways() {
	rr_sm.SidewaysMove.PrintSideways()
}

func (rr_sm *RR_sm) Random() {
	rr_sm.SidewaysMove.Cube.Random()
}

func (rr_sm *RR_sm) Run() {
	start := time.Now()

	// for RR_sm.SidewaysMove.GetValue() != 0 {
	for r := 0; r < RMAX; r++ {
		if r != 0 {
			rr_sm.Random()
		}
		// fmt.Printf("Random Restart (StA) #%d\n", r+1)
		rr_sm.SidewaysMove.Run()

		if rr_sm.SidewaysMove.GetValue() == 0 {
			break
		}
	}

	rr_sm.SidewaysMove.SetRuntime(time.Since(start))
}

func (rr_sm *RR_sm) FindBestNeighbor() *RR_sm {
	neighbor := rr_sm.SidewaysMove.FindBestNeighbor()
	return &RR_sm{SidewaysMove: *neighbor}
}
