package main

import (
    "math"
)

type Point struct {
    X, Y float64
}

func (p *Point) Rotate(a float64) {
    rad := dToR(a)
    x, y := p.X, p.Y
    p.X = x*math.Cos(rad) - y*math.Sin(rad)
    p.Y = x*math.Sin(rad) + y*math.Cos(rad)
}

func (p *Point) DistanceTo(q *Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}
