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

// verglichen mit https://www.arndt-bruenner.de/mathe/scripts/ellipsenrechner.htm
func (e Ellipse) Circumference() float64 {
    a := e.Height / 2.0
    b := e.Width / 2.0
    t := (a - b) / (a + b)
    return (a + b) * math.Pi * (1.0 + 3.0*t*t/(10.0+math.Sqrt(4.0-3.0*t*t)))
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
// Todo a = 180.0 checken
func (e Ellipse) PointInAbsoluteAngleRotated(a float64, r float64) Point {
    // "pi/2 isn't on the +Y axis" - might get strange results
    // if a == 180.0 {
    //     a = 180.0000001
    // }
    w := e.Width
    h := e.Height
    b := dToR(r)
    da := dToR(a)
    cb, sb := math.Cos(b), math.Sin(b)
    d := math.Atan2(-w*sb, h*cb)
    x := w*math.Cos(da+d)*cb - h*math.Sin(da+d)*sb // + e.X
    y := w*math.Cos(da+d)*sb + h*math.Sin(da+d)*cb // + e.Y
    return Point{x, y}
}

// https://math.stackexchange.com/questions/4086824/how-to-find-the-polar-coordinate-angle-of-the-tangent-of-any-point-on-an-ellipse/4086900#4086900
// Todo nochmal checken
func (e Ellipse) Tangent(a float64) float64 {
    return math.Atan2(e.Height*math.Cos(dToR(a)), -e.Width*math.Sin(dToR(a)))
}
