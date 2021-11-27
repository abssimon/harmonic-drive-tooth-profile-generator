package main

import (
    "math"
)

// Point

type Point struct {
    X, Y float64
}

func (p *Point) Rotate(a float64) {
    x, y := p.X, p.Y
    p.X = x*math.Cos(a) - y*math.Sin(a)
    p.Y = x*math.Sin(a) + y*math.Cos(a)
}

func (p *Point) CopyRotated(a float64) Point {
    x, y := p.X, p.Y
    pX := x*math.Cos(a) - y*math.Sin(a)
    pY := x*math.Sin(a) + y*math.Cos(a)
    return Point{pX, pY}
}

func (p *Point) DistanceTo(q *Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Cirle

type Circle struct {
    *Point // Center
    Radius float64
    Start  float64
    Stop   float64
}

func (c *Circle) Rotate(a float64) {
    c.Point.Rotate(a)
    c.Start += a
    c.Stop += a
}

func (c *Circle) PointInAngle(a float64) Point {
    x := math.Cos(a) * c.Radius
    y := math.Sin(a) * c.Radius
    return Point{x + c.X, y + c.Y}
}

// https://www.weltderfertigung.de/suchen/lernen/mathematik/beruehrpunktberechnung-tangente-an-zwei-kreisen.php
func (c *Circle) InnerTangentWith(c2 *Circle) (float64, float64) {
    hp := c.Point.DistanceTo(c2.Point)
    v := c2.Radius / c.Radius
    len1 := hp / 100 * (100.0 / (v + 1.0) * v)

    // angles
    x := c2.Radius / len1
    var a1 float64
    if x >= 1.0 {
        a1 = math.Pi - (math.Sin(x) + math.Pi/2)
    } else {
        a1 = math.Pi - (math.Asin(x) + math.Pi/2)
    }

    a2 := math.Asin((c.Y - c2.Y) / hp)
    a3 := math.Pi/2 - (a2 + a1)
    a4 := a1 - a2

    return a3, a4
}

func (c *Circle) Coordinates() []Point {
    points := []Point{}
    for i := c.Start; i <= c.Stop; i += 0.03491 { // 2 degree resolution
        points = append(points, c.PointInAngle(i))
    }
    points = append(points, c.PointInAngle(c.Stop))

    return points
}

func newCircle(x, y, r float64) *Circle {
    return &Circle{&Point{x, y}, r, 0.0, 0.0}
}

// Ellipse

type Ellipse struct {
    Point
    Height float64
    Width  float64
}

// https://www.arndt-bruenner.de/mathe/scripts/ellipsenrechner.htm
func (e Ellipse) Circumference() float64 {
    a := e.Height / 2.0
    b := e.Width / 2.0
    t := (a - b) / (a + b)
    return (a + b) * math.Pi * (1.0 + 3.0*t*t/(10.0+math.Sqrt(4.0-3.0*t*t)))
}

func (e Ellipse) PointAtAngle(a float64) Point {
    return Point{e.Width * math.Cos(a), e.Height * math.Sin(a)}
}

// https://math.stackexchange.com/questions/4086824
func (e Ellipse) Tangent(a float64) float64 {
    return math.Atan2(e.Height*math.Cos(a), -e.Width*math.Sin(a))
}

// TangentByPoint is used to calculate a tangens for a rotated ellipse.
// A point from a rotated ellipse can be rotated back, and with this
// a normal tangens can be calculated
// https://math.stackexchange.com/questions/4134884
func (e Ellipse) TangentByPoint(p Point, i float64) float64 {
    a := math.Atan2(p.Y/e.Height, p.X/e.Width)
    return math.Atan2(e.Height*math.Cos(a), -e.Width*math.Sin(a))
}
