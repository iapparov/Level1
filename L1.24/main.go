package main

import (
	"fmt"
	"math"
)


type Point struct{
	x,y float64
}

func NewPoint(x,y float64) *Point{
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(o *Point) float64{
	return math.Sqrt(math.Pow(o.x-p.x,2)+math.Pow(o.y-p.y,2))
}
func main() {
	p1 := NewPoint(1.5, 2.5)
	p2 := NewPoint(2.3, 3.7)
	distance := p1.Distance(p2)
	fmt.Println(distance)
}