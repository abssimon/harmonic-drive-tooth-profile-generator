package main

import (
    "math"
    "errors"
)

type Point struct {
    X float64
    Y float64
}

func (p *Point) DistanceTo(q *Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) Rotate(a float64) {
    x, y := p.X, p.Y
    p.X = x*math.Cos(a) - y*math.Sin(a)
    p.Y = x*math.Sin(a) + y*math.Cos(a)
}

// https://github.com/psalaets/line-intersect/blob/master/src/check-intersection.js
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

// ---

type Circle struct {
    *Point
    Radius float64
    Start  float64
    Stop   float64
}

func (c *Circle) PointInAngle(a float64) Point {
    x := math.Cos(a) * c.Radius
    y := math.Sin(a) * c.Radius
    return Point{x + c.X, y + c.Y}
}

func (c *Circle) Coordinates() []Point {
    num := int(math.Abs((c.Start - c.Stop) / 0.03491)) // 2 deg. circle resolution for lower file size
    r := make([]Point, num)
    p := LinSpace(c.Start, c.Stop, num)
    for i, v := range p {
        r[i] = c.PointInAngle(v)
    }
    return r
}

func (c *Circle) Rotate(a float64) {
    c.Point.Rotate(a)
    c.Start += a
    c.Stop += a
}

// https://www.weltderfertigung.de/suchen/lernen/mathematik/beruehrpunktberechnung-tangente-an-zwei-kreisen.php
func (c *Circle) InnerTangentWith(c2 *Circle) (float64, float64) {
    hp := c.Point.DistanceTo(c2.Point)
    v := c2.Radius / c.Radius
    len1 := hp / 100 * (100.0 / (v + 1.0) * v)

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

// ---

type Ellipse struct {
    Point
    Height float64
    Width  float64
}

func (e Ellipse) Coordinates(num int) []Point {
    r := make([]Point, num)
    p := LinSpace(0.0, 2.0*math.Pi, num)
    for i, v := range p {
        r[i] = Point{math.Cos(v) * e.Height / 2.0, math.Sin(v) * e.Width / 2.0}
    }
    return r
}

// https://math.stackexchange.com/questions/4134884
func (e Ellipse) TangentByPoint(p Point) float64 {
    a := math.Atan2(p.Y/e.Width, p.X/e.Height)
    return math.Atan2(e.Width*math.Cos(a), -e.Height*math.Sin(a))
}
