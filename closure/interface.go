package main

import (
	"fmt"
	"math"
)

type Describer interface {
	describe()
}

type Circle struct {
	radius float64
}

type Cylinder struct {
	radius float64
	height float64
}

func (c Circle) describe() {
	area := math.Pi * math.Pow(c.radius, 2)
	circ := 2 * math.Pi * c.radius
	fmt.Printf("A circle with radius %v,\narea %v,\nand circumference %v\n\n", c.radius, area, circ)
}

func (c Cylinder) describe() {
	volume := math.Pi * math.Pow(c.radius, 2) * c.height
	circ := 2 * math.Pi * c.radius
	fmt.Printf("A cylinder with radius %v,\nvolume %v,\nand circumference %v\n\n", c.radius, volume, circ)
}

func main() {
	circle := Circle{5}
	cylinder := Cylinder{5, 3}
	circle.describe()
	cylinder.describe()
}
