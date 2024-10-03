package main

import (
	"time"
)

const RMAX = 25

type RR_sta struct {
	SteepestAscent
}

func NewRR_sta(cube *Cube) *RR_sta {
	sta := &SteepestAscent{
		Cube: *cube,
	}

	return &RR_sta{
		SteepestAscent: *sta,
	}
}

func (rr_sta *RR_sta) GetConfiguration() [DIMENSION][DIMENSION][DIMENSION]uint8 {
	return rr_sta.SteepestAscent.GetConfiguration()
}

func (rr_sta *RR_sta) GetValue() int {
	return rr_sta.SteepestAscent.GetValue()
}

func (rr_sta *RR_sta) GetValue1() int {
	return rr_sta.SteepestAscent.GetValue1()
}

func (rr_sta *RR_sta) GetValue2() int {
	return rr_sta.SteepestAscent.GetValue2()
}

func (rr_sta *RR_sta) GetCube() *Cube {
	return rr_sta.SteepestAscent.GetCube()
}

func (rr_sta *RR_sta) GetRuntime() time.Duration {
	return rr_sta.SteepestAscent.GetRuntime()
}

func (rr_sta *RR_sta) SetValue() {
	rr_sta.SteepestAscent.SetValue()
}

func (rr_sta *RR_sta) SetRuntime(runtime time.Duration) {
	rr_sta.Runtime = runtime
}

func (rr_sta *RR_sta) Clone() *RR_sta {
	return &RR_sta{SteepestAscent: *rr_sta.SteepestAscent.Clone()}
}

func (rr_sta *RR_sta) Copy(original *RR_sta) {
	rr_sta.SteepestAscent.Copy(&original.SteepestAscent)
}

func (rr_sta *RR_sta) PrintSideways() {
	rr_sta.SteepestAscent.PrintSideways()
}

func (rr_sta *RR_sta) Random() {
	rr_sta.SteepestAscent.Cube.Random()
}

func (rr_sta *RR_sta) Run() {
	start := time.Now()

	// for rr_sta.SteepestAscent.GetValue() != 0 {
	for r := 0; r < RMAX; r++ {
		if r != 0 {
			rr_sta.Random()
		}
		// fmt.Printf("Random Restart (StA) #%d\n", r+1)
		rr_sta.SteepestAscent.Run()

		if rr_sta.SteepestAscent.GetValue() == 0 {
			break
		}
	}

	rr_sta.SteepestAscent.Runtime = time.Since(start)
}

func (rr_sta *RR_sta) FindBestNeighbor() *RR_sta {
	neighbor := rr_sta.SteepestAscent.FindBestNeighbor()
	return &RR_sta{SteepestAscent: *neighbor}
}
