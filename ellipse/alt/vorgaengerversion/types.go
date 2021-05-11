package main

import (
    "math"
)

type Point struct {
    X, Y float64
}

type Circle struct {
    Point
    Radius float64
}

func (c Circle) PointInAngle(a float64) Point {
    x := math.Cos(dToR(a))*c.Radius + c.X
    y := math.Sin(dToR(a))*c.Radius + c.Y
    return Point{x, y}
}

func newCircle(x, y, r float64) *Circle {
    return &Circle{Point{x, y}, r}
}

type Ellipse struct {
    Point
    Height float64
    Width  float64
}

func (e Ellipse) PointInAngle(a float64) Point {
    da := dToR(a)
    x := e.Width*math.Cos(da) + e.X
    y := e.Height*math.Sin(da) + e.Y
    return Point{x, y}
}

// https://math.stackexchange.com/questions/2645689/what-is-the-parametric-equation-of-a-rotated-ellipse-given-the-angle-of-rotatio
func (e Ellipse) PointInAngleRotated(a float64, r float64) Point {
    da := dToR(a)
    dr := dToR(r)
    x := e.Width*math.Cos(da)*math.Cos(dr) - e.Height*math.Sin(da)*math.Sin(dr) + e.X // optimise...
    y := e.Width*math.Cos(da)*math.Sin(dr) + e.Height*math.Sin(da)*math.Cos(dr) + e.Y
    return Point{x, y}
}

// https://math.stackexchange.com/questions/4130633/rotated-ellipse-calculate-points-with-an-absolute-angle
func (e Ellipse) PointInAbsoluteAngleRotated(aa float64, r float64) Point {
    w := e.Width
    h := e.Height
    b := dToR(r)
    a := dToR(aa)
    cb, sb := math.Cos(b), math.Sin(b)
    d := math.Atan2(-w*sb, h*cb)
    x := w*math.Cos(a+d)*cb - h*math.Sin(a+d)*sb + e.X
    y := w*math.Cos(a+d)*sb + h*math.Sin(a+d)*cb + e.Y
    return Point{x, y}
}

func (e Ellipse) Circumference() float64 {
    a := e.Height / 2.0
    b := e.Width / 2.0
    t := (a - b) / (a + b)
    return (a + b) * math.Pi * (1.0 + 3.0*t*t/(10.0+math.Sqrt(4.0-3.0*t*t)))
}
