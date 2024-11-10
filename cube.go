package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	DIMENSION    = 5
	MAGIC_NUMBER = 315
	ELEMENT      = DIMENSION * DIMENSION * DIMENSION
)

type Cube struct {
	Dimension     uint8
	Configuration [DIMENSION][DIMENSION][DIMENSION]uint8
	Value         int
	Value1        int
	Value2        int
}

func NewCube() *Cube {
	cube := &Cube{Dimension: DIMENSION}
	cube.Random()

	return cube
}

// Objective function for the cube
func (c *Cube) evalValue(value *int) {
	if (*value) != MAGIC_NUMBER {
		c.Value++
	}
	c.Value1 += int(math.Abs(float64(int(*value) - MAGIC_NUMBER)))
	c.Value2 += int(float64((int(*value) - MAGIC_NUMBER) * (int(*value) - MAGIC_NUMBER)))
}

func (c *Cube) SetValue() {
	c.Value = 0
	c.Value1 = 0
	c.Value2 = 0
	counter := uint8(0)
	tmp := int(0)

	// Evaluate the sum of each row in each layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			for k := uint8(0); k < c.Dimension; k++ {
				tmp += int(c.Configuration[i][j][k])
			}
			// fmt.Printf("Sum of row (%d, %d): %d\n", i, j, tmp)
			c.evalValue(&tmp)
			tmp = 0
			counter++
		}
	}

	// Evaluate the sum of each column in each layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			for k := uint8(0); k < c.Dimension; k++ {
				tmp += int(c.Configuration[i][k][j])
			}
			// fmt.Printf("Sum of column (%d, %d): %d\n", i, j, tmp)
			c.evalValue(&tmp)
			tmp = 0
			counter++
		}
	}

	// Evaluate the sum of each pillars in cube
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			for k := uint8(0); k < c.Dimension; k++ {
				tmp += int(c.Configuration[k][i][j])
			}
			// fmt.Printf("Sum of pillar (%d, %d): %d\n", i, j, tmp)
			c.evalValue(&tmp)
			tmp = 0
			counter++
		}
	}

	// Evaluate the value of diagonal in front facing layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			tmp += int(c.Configuration[i][j][j])
		}
		// fmt.Printf("Sum of diagonal (%d): %d\n", i, tmp)
		c.evalValue(&tmp)
		tmp = 0
		counter++
	}

	// Evaluate the value of diagonal in back facing layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			tmp += int(c.Configuration[i][j][c.Dimension-1-j])
		}
		// fmt.Printf("Sum of diagonal (%d): %d\n", i, tmp)
		c.evalValue(&tmp)
		tmp = 0
		counter++
	}

	// Evaluate the value of diagonal in left facing layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			tmp += int(c.Configuration[j][i][j])
		}
		// fmt.Printf("Sum of diagonal (%d): %d\n", i, tmp)
		c.evalValue(&tmp)
		tmp = 0
		counter++
	}

	// Evaluate the value of diagonal in right facing layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			tmp += int(c.Configuration[j][i][c.Dimension-1-j])
		}
		// fmt.Printf("Sum of diagonal (%d): %d\n", i, tmp)
		c.evalValue(&tmp)
		tmp = 0
		counter++
	}

	// Evaluate the value of diagonal in top facing layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			tmp += int(c.Configuration[j][j][i])
		}
		// fmt.Printf("Sum of diagonal (%d): %d\n", i, tmp)
		c.evalValue(&tmp)
		tmp = 0
		counter++
	}

	// Evaluate the value of diagonal in bottom facing layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			tmp += int(c.Configuration[j][c.Dimension-1-j][i])
		}
		// fmt.Printf("Sum of diagonal (%d): %d\n", i, tmp)
		c.evalValue(&tmp)
		tmp = 0
		counter++
	}

	// Evaluate the value of triagonal in cube (from top-left to bottom-right)
	for i := uint8(0); i < c.Dimension; i++ {
		tmp += int(c.Configuration[i][i][i])
	}
	// fmt.Printf("Sum of triagonal (top-left to bottom-right): %d\n", tmp)
	c.evalValue(&tmp)
	tmp = 0
	counter++

	// Evaluate the value of triagonal in cube (from top-right to bottom-left)
	for i := uint8(0); i < c.Dimension; i++ {
		tmp += int(c.Configuration[i][c.Dimension-1-i][i])
	}
	// fmt.Printf("Sum of triagonal (top-right to bottom-left): %d\n", tmp)
	c.evalValue(&tmp)
	tmp = 0
	counter++

	// Evaluate the value of triagonal in cube (from bottom-left to top-right)
	for i := uint8(0); i < c.Dimension; i++ {
		tmp += int(c.Configuration[c.Dimension-1-i][i][i])
	}
	// fmt.Printf("Sum of triagonal (bottom-left to top-right): %d\n", tmp)
	c.evalValue(&tmp)
	tmp = 0
	counter++

	// Evaluate the value of triagonal in cube (from bottom-right to top-left)
	for i := uint8(0); i < c.Dimension; i++ {
		tmp += int(c.Configuration[c.Dimension-1-i][c.Dimension-1-i][i])
	}
	// fmt.Printf("Sum of triagonal (bottom-right to top-left): %d\n", tmp)
	c.evalValue(&tmp)
	tmp = 0
	counter++
}

// Generate random initial configuration for the cube
func (c *Cube) Random() {
	// Create a slice with values 1 to 125
	slice := make([]uint8, 125)
	for i := range slice {
		slice[i] = uint8(i + 1)
	}

	// Shuffle the slice
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	// Assign shuffled values to the 3D configuration of the cube
	c.unflatten(slice)
}

func (c *Cube) PrintSideways() {
	for i := uint8(0); i < c.Dimension; i++ {
		for z := uint8(0); z < c.Dimension; z++ {
			for j := uint8(0); j < c.Dimension; j++ {
				fmt.Printf("%3d ", c.Configuration[z][i][j])
			}
			fmt.Printf("\t")
		}
		fmt.Println()
	}
}

func (c *Cube) FindBestNeighbor() *Cube {
	flat := c.flatten()
	unflat := c.Clone()
	best := c.Clone()

	for i := 0; i < ELEMENT-1; i++ {
		for j := i + 1; j < ELEMENT; j++ {
			flat[i], flat[j] = flat[j], flat[i]

			unflat.unflatten(flat)
			if unflat.Value < best.Value {
				best.Copy(unflat)
			}

			flat[i], flat[j] = flat[j], flat[i]
		}
	}

	return best
}

func (c *Cube) FindRandomNeighbor() {
	x1 := uint8(rand.Intn(5))
	y1 := uint8(rand.Intn(5))
	z1 := uint8(rand.Intn(5))
	x2, y2, z2 := x1, y1, z1

	for x1 == x2 && y1 == y2 && z1 == z2 {
		x2 = uint8(rand.Intn(5))
		y2 = uint8(rand.Intn(5))
		z2 = uint8(rand.Intn(5))
	}

	c.Swap(x1, y1, z1, x2, y2, z2)
}

// // Print the cube configuration
// func (c *Cube) PrintConfiguration() {
// 	for i := uint8(0); i < c.Dimension; i++ {
// 		for j := uint8(0); j < c.Dimension; j++ {
// 			for k := uint8(0); k < c.Dimension; k++ {
// 				fmt.Printf("%d ", c.Configuration[i][j][k])
// 			}
// 			fmt.Println()
// 		}
// 		fmt.Println()
// 	}
// }

func (c *Cube) Copy(original *Cube) {
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			for k := uint8(0); k < c.Dimension; k++ {
				c.Configuration[i][j][k] = original.Configuration[i][j][k]
			}
		}
	}
	c.Value = original.Value
	c.Value1 = original.Value1
	c.Value2 = original.Value2
}

func (c *Cube) Clone() *Cube {
	newCube := &Cube{
		Dimension: c.Dimension,
		Value:     c.Value,
		Value1:    c.Value1,
		Value2:    c.Value2,
	}

	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			for k := uint8(0); k < c.Dimension; k++ {
				newCube.Configuration[i][j][k] = c.Configuration[i][j][k]
			}
		}
	}
	return newCube
}

func (c *Cube) Swap(x1, y1, z1, x2, y2, z2 uint8) {
	c.Configuration[x1][y1][z1], c.Configuration[x2][y2][z2] = c.Configuration[x2][y2][z2], c.Configuration[x1][y1][z1]
	c.SetValue()
}

// func FlatSwap(flat *[]uint8, i, j int) {
// 	(*flat)[i], (*flat)[j] = (*flat)[j], (*flat)[i]
// }

func (c *Cube) flatten() []uint8 {
	flat := make([]uint8, 0, c.Dimension*c.Dimension*c.Dimension)
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			for k := uint8(0); k < c.Dimension; k++ {
				flat = append(flat, c.Configuration[i][j][k])
			}
		}
	}
	return flat
}

func (c *Cube) unflatten(flat []uint8) {
	idx := 0
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			for k := uint8(0); k < c.Dimension; k++ {
				c.Configuration[i][j][k] = flat[idx]
				idx++
			}
		}
	}

	c.SetValue()
}

func (c1 *Cube) IsSame(c2 *Cube) bool {
	for i := uint8(0); i < c1.Dimension; i++ {
		for j := uint8(0); j < c1.Dimension; j++ {
			for k := uint8(0); k < c1.Dimension; k++ {
				if c1.Configuration[i][j][k] != c2.Configuration[i][j][k] {
					return false
				}
			}
		}
	}
	return true
}

func (c *Cube) GetDimension() uint8 {
	return c.Dimension
}

func (c *Cube) GetValue() int {
	return c.Value
}

func (c *Cube) GetValue1() int {
	return c.Value1
}

func (c *Cube) GetValue2() int {
	return c.Value2
}

func (c *Cube) GetConfiguration() [DIMENSION][DIMENSION][DIMENSION]uint8 {
	return c.Configuration
}
