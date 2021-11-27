package main

import (
	"errors"
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

func intersection(p1, p2, p3, p4 Point) (Point, error) {
	denom := ((p4.Y - p3.Y) * (p2.X - p1.X)) - ((p4.X - p3.X) * (p2.Y - p1.Y))
	numeA := ((p4.X - p3.X) * (p1.Y - p3.Y)) - ((p4.Y - p3.Y) * (p1.X - p3.X))
	numeB := ((p2.X - p1.X) * (p1.Y - p3.Y)) - ((p2.Y - p1.Y) * (p1.X - p3.X))

	if denom == 0.0 {
		if numeA == 0.0 && numeB == 0.0 {
			return Point{}, errors.New("COLINEAR")
		}
		return Point{}, errors.New("PARALLEL")
	}
	uA := numeA / denom
	uB := numeB / denom

	if uA >= 0.0 && uA <= 1.0 && uB >= 0.0 && uB <= 1.0 {
		x := p1.X + (uA * (p2.X - p1.X))
		y := p1.Y + (uA * (p2.Y - p1.Y))
		return Point{x, y}, nil
	}

	return Point{}, errors.New("NONE")
}

func main() {

	p1 := Point{3.6, -1.0}
	p2 := Point{4.05, -2.8}

	p3 := Point{1.9, -2.9}
	p4 := Point{3.6, -3.6}

	// verlaengern

	p2.X += p2.X - p1.X
	p2.Y += p2.Y - p1.Y

	p4.X += p4.X - p3.X
	p4.Y += p4.Y - p3.Y

	fmt.Println(intersection(p1, p2, p3, p4))

}
