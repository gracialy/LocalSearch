package main

import (
	"time"
)

type SteepestAscent struct {
	Cube
	Runtime time.Duration
}

func NewSteepestAscent(cube *Cube) *SteepestAscent {
	sta := &SteepestAscent{
		Cube: *cube,
	}

	sta.Runtime = 0

	return sta
}

func (sta *SteepestAscent) GetConfiguration() [DIMENSION][DIMENSION][DIMENSION]uint8 {
	return sta.Cube.GetConfiguration()
}

func (sta *SteepestAscent) GetValue() int {
	return sta.Cube.GetValue()
}

func (sta *SteepestAscent) GetValue1() int {
	return sta.Cube.GetValue1()
}

func (sta *SteepestAscent) GetValue2() int {
	return sta.Cube.GetValue2()
}

func (sta *SteepestAscent) GetCube() *Cube {
	return &sta.Cube
}

func (sta *SteepestAscent) GetRuntime() time.Duration {
	return sta.Runtime
}

func (sta *SteepestAscent) SetValue() {
	sta.Cube.SetValue()
}

func (sta *SteepestAscent) SetRuntime(runtime time.Duration) {
	sta.Runtime = runtime
}

func (sta *SteepestAscent) Clone() *SteepestAscent {
	return &SteepestAscent{Cube: *sta.Cube.Clone()}
}

func (sta *SteepestAscent) Copy(original *SteepestAscent) {
	sta.Cube.Copy(original.GetCube())
	sta.Runtime = original.GetRuntime()
}

func (sta *SteepestAscent) PrintSideways() {
	sta.Cube.PrintSideways()
}

func (sta *SteepestAscent) Run() {
	start := time.Now()

	current := sta

	for {
		neighbor := sta.FindBestNeighbor()
		if neighbor.GetValue() < current.GetValue() {
			current.Copy(neighbor)
		} else {
			break
		}
	}

	sta.Runtime = time.Since(start)
}

func (sta *SteepestAscent) FindBestNeighbor() *SteepestAscent {
	flat := sta.flatten()
	unflat := sta.Clone()
	best := sta.Clone()

	for i := 0; i < ELEMENT-1; i++ {
		for j := i + 1; j < ELEMENT; j++ {
			flat[i], flat[j] = flat[j], flat[i]

			unflat.unflatten(flat)
			if unflat.GetValue() < best.GetValue() {
				best.Copy(unflat)
			}

			flat[i], flat[j] = flat[j], flat[i]
		}
	}

	return best
}

func (sta *SteepestAscent) flatten() []uint8 {
	return sta.Cube.flatten()
}

func (sta *SteepestAscent) unflatten(flat []uint8) {
	sta.Cube.unflatten(flat)
}
