package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	DIMENSION    = 5
	MAGIC_NUMBER = 315
)

type Cube struct {
	Dimension     uint8
	Configuration [DIMENSION][DIMENSION][DIMENSION]uint8
	Value         uint8
}

func NewCube() *Cube {
	cube := &Cube{Dimension: DIMENSION}
	cube.Random()
	cube.SetValue()

	return cube
}

func (c *Cube) GetDimension() uint8 {
	return c.Dimension
}

func (c *Cube) GetValue() uint8 {
	return c.Value
}

func (c *Cube) GetConfiguration() [DIMENSION][DIMENSION][DIMENSION]uint8 {
	return c.Configuration
}

func (c *Cube) SetValue() {
	c.Value = 0
	counter := uint8(0)
	tmp := uint16(0)

	// Evaluate the sum of each row in each layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			for k := uint8(0); k < c.Dimension; k++ {
				tmp += uint16(c.Configuration[i][j][k])
			}
			// fmt.Printf("Sum of row (%d, %d): %d\n", i, j, tmp)
			c.addValue(&tmp)
			tmp = 0
			counter++
		}
	}

	// Evaluate the sum of each column in each layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			for k := uint8(0); k < c.Dimension; k++ {
				tmp += uint16(c.Configuration[i][k][j])
			}
			// fmt.Printf("Sum of column (%d, %d): %d\n", i, j, tmp)
			c.addValue(&tmp)
			tmp = 0
			counter++
		}
	}

	// Evaluate the sum of each pillars in cube
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			for k := uint8(0); k < c.Dimension; k++ {
				tmp += uint16(c.Configuration[k][i][j])
			}
			// fmt.Printf("Sum of pillar (%d, %d): %d\n", i, j, tmp)
			c.addValue(&tmp)
			tmp = 0
			counter++
		}
	}

	// Evaluate the value of diagonal in front facing layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			tmp += uint16(c.Configuration[i][j][j])
		}
		// fmt.Printf("Sum of diagonal (%d): %d\n", i, tmp)
		c.addValue(&tmp)
		tmp = 0
		counter++
	}

	// Evaluate the value of diagonal in back facing layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			tmp += uint16(c.Configuration[i][j][c.Dimension-1-j])
		}
		// fmt.Printf("Sum of diagonal (%d): %d\n", i, tmp)
		c.addValue(&tmp)
		tmp = 0
		counter++
	}

	// Evaluate the value of diagonal in left facing layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			tmp += uint16(c.Configuration[j][i][j])
		}
		// fmt.Printf("Sum of diagonal (%d): %d\n", i, tmp)
		c.addValue(&tmp)
		tmp = 0
		counter++
	}

	// Evaluate the value of diagonal in right facing layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			tmp += uint16(c.Configuration[j][i][c.Dimension-1-j])
		}
		// fmt.Printf("Sum of diagonal (%d): %d\n", i, tmp)
		c.addValue(&tmp)
		tmp = 0
		counter++
	}

	// Evaluate the value of diagonal in top facing layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			tmp += uint16(c.Configuration[j][j][i])
		}
		// fmt.Printf("Sum of diagonal (%d): %d\n", i, tmp)
		c.addValue(&tmp)
		tmp = 0
		counter++
	}

	// Evaluate the value of diagonal in bottom facing layer
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			tmp += uint16(c.Configuration[j][c.Dimension-1-j][i])
		}
		// fmt.Printf("Sum of diagonal (%d): %d\n", i, tmp)
		c.addValue(&tmp)
		tmp = 0
		counter++
	}

	// Evaluate the value of triagonal in cube (from top-left to bottom-right)
	for i := uint8(0); i < c.Dimension; i++ {
		tmp += uint16(c.Configuration[i][i][i])
	}
	// fmt.Printf("Sum of triagonal (top-left to bottom-right): %d\n", tmp)
	c.addValue(&tmp)
	tmp = 0
	counter++

	// Evaluate the value of triagonal in cube (from top-right to bottom-left)
	for i := uint8(0); i < c.Dimension; i++ {
		tmp += uint16(c.Configuration[i][c.Dimension-1-i][i])
	}
	// fmt.Printf("Sum of triagonal (top-right to bottom-left): %d\n", tmp)
	c.addValue(&tmp)
	tmp = 0
	counter++

	// Evaluate the value of triagonal in cube (from bottom-left to top-right)
	for i := uint8(0); i < c.Dimension; i++ {
		tmp += uint16(c.Configuration[c.Dimension-1-i][i][i])
	}
	// fmt.Printf("Sum of triagonal (bottom-left to top-right): %d\n", tmp)
	c.addValue(&tmp)
	tmp = 0
	counter++

	// Evaluate the value of triagonal in cube (from bottom-right to top-left)
	for i := uint8(0); i < c.Dimension; i++ {
		tmp += uint16(c.Configuration[c.Dimension-1-i][c.Dimension-1-i][i])
	}
	// fmt.Printf("Sum of triagonal (bottom-right to top-left): %d\n", tmp)
	c.addValue(&tmp)
	tmp = 0
	counter++
}

// Generate random initial configuration for the cube
func (c *Cube) Random() {
	// Create a slice with values 1 to 125
	values := make([]uint8, 125)
	for i := range values {
		values[i] = uint8(i + 1)
	}

	// Shuffle the slice
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	// Assign shuffled values to the 3D configuration of the cube
	idx := 0
	for i := 0; i < int(c.Dimension); i++ {
		for j := 0; j < int(c.Dimension); j++ {
			for k := 0; k < int(c.Dimension); k++ {
				c.Configuration[i][j][k] = values[idx]
				idx++
			}
		}
	}
}

// Print the cube configuration
func (c *Cube) PrintConfiguration() {
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			for k := uint8(0); k < c.Dimension; k++ {
				fmt.Printf("%d ", c.Configuration[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func (c *Cube) Copy(original *Cube) {
	for i := uint8(0); i < c.Dimension; i++ {
		for j := uint8(0); j < c.Dimension; j++ {
			for k := uint8(0); k < c.Dimension; k++ {
				c.Configuration[i][j][k] = original.Configuration[i][j][k]
			}
		}
	}
	c.Value = original.Value
}

func (c *Cube) Clone() *Cube {
	newCube := &Cube{
		Dimension: c.Dimension,
		Value:     c.Value,
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

func (c *Cube) addValue(value *uint16) {
	if (*value) == MAGIC_NUMBER {
		c.Value++
	}
}
