package main

import (
	"time"
)

type SidewaysMove struct {
	SteepestAscent
}

func NewSidewaysMove(cube *Cube) *SidewaysMove {
	sa := &SteepestAscent{
		Cube: *cube,
	}

	return &SidewaysMove{
		SteepestAscent: *sa,
	}
}

func (sm *SidewaysMove) GetConfiguration() [DIMENSION][DIMENSION][DIMENSION]uint8 {
	return sm.SteepestAscent.GetConfiguration()
}

func (sm *SidewaysMove) GetValue() int {
	return sm.SteepestAscent.GetValue()
}

func (sm *SidewaysMove) GetValue1() int {
	return sm.SteepestAscent.GetValue1()
}

func (sm *SidewaysMove) GetValue2() int {
	return sm.SteepestAscent.GetValue2()
}

func (sm *SidewaysMove) GetCube() *Cube {
	return sm.SteepestAscent.GetCube()
}

func (sm *SidewaysMove) GetRuntime() time.Duration {
	return sm.SteepestAscent.GetRuntime()
}

func (sm *SidewaysMove) SetValue() {
	sm.SteepestAscent.SetValue()
}

func (sm *SidewaysMove) SetRuntime(runtime time.Duration) {
	sm.SteepestAscent.SetRuntime(runtime)
}

func (sm *SidewaysMove) Clone() *SidewaysMove {
	return &SidewaysMove{SteepestAscent: *sm.SteepestAscent.Clone()}
}

func (sm *SidewaysMove) Copy(original *SidewaysMove) {
	sm.SteepestAscent.Copy(&original.SteepestAscent)
}

func (sm *SidewaysMove) PrintSideways() {
	sm.SteepestAscent.PrintSideways()
}

func (sm *SidewaysMove) Run() {
	start := time.Now()

	current := sm
	stuck := 0

	for {
		neighbor := sm.FindBestNeighbor()
		if neighbor.GetValue() > current.GetValue() {
			break
		} else if neighbor.GetValue() == current.GetValue() {
			stuck++
		} else {
			stuck = 0
		}

		current.Copy(neighbor)

		if stuck == 100 {
			break
		}
	}

	sm.SteepestAscent.Runtime = time.Since(start)
}

func (sm *SidewaysMove) FindBestNeighbor() *SidewaysMove {
	neighbor := sm.SteepestAscent.FindBestNeighbor()
	return &SidewaysMove{SteepestAscent: *neighbor}
}
