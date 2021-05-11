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

func (p *Point) CopyRotated(a float64) Point {
    rad := dToR(a)
    x, y := p.X, p.Y
    pX := x*math.Cos(rad) - y*math.Sin(rad)
    pY := x*math.Sin(rad) + y*math.Cos(rad)
    return Point{pX, pY}
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

// same as https://www.arndt-bruenner.de/mathe/scripts/ellipsenrechner.htm
func (e Ellipse) Circumference() float64 {
    a := e.Height / 2.0
    b := e.Width / 2.0
    t := (a - b) / (a + b)
    return (a + b) * math.Pi * (1.0 + 3.0*t*t/(10.0+math.Sqrt(4.0-3.0*t*t)))
}

func (e Ellipse) PointInAngle(a float64) Point {
    da := dToR(a)
    x := e.Width*math.Cos(da) 
    y := e.Height*math.Sin(da)
    return Point{x, y}
}

// https://math.stackexchange.com/questions/2645689/what-is-the-parametric-equation-of-a-rotated-ellipse-given-the-angle-of-rotatio
func (e Ellipse) PointByAngleRotated(a float64, r float64) Point {
    da := dToR(a)
    dr := dToR(r)
    x := e.Width*math.Cos(da)*math.Cos(dr) - e.Height*math.Sin(da)*math.Sin(dr) 
    y := e.Width*math.Cos(da)*math.Sin(dr) + e.Height*math.Sin(da)*math.Cos(dr) 
    return Point{x, y}
}

// PointInAbsoluteAngleRotated is used to fix a point (teeth), so it wont
// rotate with the elipse.
// https://math.stackexchange.com/questions/4130633/rotated-ellipse-calculate-points-with-an-absolute-angle
// TODO a = 180.0 checken "pi/2 isn't on the +Y axis"
func (e Ellipse) PointByAbsoluteAngleRotated(a float64, r float64) Point {
    panic("Not round")
    w := e.Width
    h := e.Height
    b := dToR(r)
    da := dToR(a)
    cb, sb := math.Cos(b), math.Sin(b)
    d := math.Atan2(-w*sb, h*cb)
    x := w*math.Cos(da+d)*cb - h*math.Sin(da+d)*sb 
    y := w*math.Cos(da+d)*sb + h*math.Sin(da+d)*cb 
    return Point{x, y}
}

// https://math.stackexchange.com/questions/4086824/how-to-find-the-polar-coordinate-angle-of-the-tangent-of-any-point-on-an-ellipse/4086900#4086900
func (e Ellipse) Tangent(a float64) float64 {
    return math.Atan2(e.Height*math.Cos(dToR(a)), -e.Width*math.Sin(dToR(a)))
}

// TangentByPoint is used to calculate a tangens for a rotated ellipse.
// A point from a rotated ellipse can be rotated back, and with this 
// a normal tangens can be calculated
func (e Ellipse) TangentByPoint(p Point, i float64) float64 {
    a := math.Atan2(p.Y/e.Height, p.X/e.Width) 
    return math.Atan2(e.Height*math.Cos(a), -e.Width*math.Sin(a))
}
