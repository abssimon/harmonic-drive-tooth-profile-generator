package main

import (
    "math"
)

type Ellipse struct {
    *Point
    Width  float64
    Height float64
}

func (e Ellipse) PointInAngle(a float64) Point {
    x := e.Width*math.Cos(dToR(a)) + e.X
    y := e.Height*math.Sin(dToR(a)) + e.Y
    return Point{x, y}
}

func newEllipse(x, y, w, h float64) *Ellipse {
    return &Ellipse{&Point{x, y}, w, h}
}
