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

type Line struct {
    p1 *Point
    p2 *Point
}

type Circle struct {
    *Point  // Center
    Radius float64
    StartPoint *Point
    StopPoint *Point
}

func (c *Circle) Rotate(a float64) {
    c.Point.Rotate(a)
    c.StartPoint.Rotate(a)
    c.StopPoint.Rotate(a)
}


func (c *Circle) PointInAngle(a float64) Point {
    rad := dToR(a)
    x := math.Cos(rad) * c.Radius
    y := math.Sin(rad) * c.Radius

    return Point{x + c.X, y + c.Y}
}


// https://www.weltderfertigung.de/suchen/lernen/mathematik/beruehrpunktberechnung-tangente-an-zwei-kreisen.php
func (c *Circle) InnerTangentTo(c2 *Circle) (Line, Line) {
    hp := c.Point.DistanceTo(c2.Point)
    v := c2.Radius / c.Radius
    len1 := hp / 100 * (100.0 / (v + 1.0) * v)

    // angles
    x := c2.Radius / len1
    var a1 float64
    if x >= 1.0 {
        a1 = 180.0 - (rToD(math.Sin(x)) + 90.0)
    } else {
        a1 = 180.0 - (rToD(math.Asin(x)) + 90.0)
    }
    a2 := rToD(math.Asin((c.Y - c2.Y) / hp))
    a3 := 90 - (a2 + a1)
    a4 := a1 - a2

    rad1 := dToR(a3)
    rad2 := dToR(a4)

    sa3 := math.Sin(rad1) * c2.Radius // todo names
    sb3 := math.Cos(rad1) * c2.Radius
    sa5 := math.Sin(rad2) * c2.Radius
    sb5 := math.Cos(rad2) * c2.Radius

    osa5 := math.Sin(rad1) * c.Radius
    osb5 := math.Cos(rad1) * c.Radius
    osa7 := math.Sin(rad2) * c.Radius
    osb7 := math.Cos(rad2) * c.Radius

    return Line{&Point{c2.X - sa3, c2.Y + sb3}, &Point{c.X + osa5, c.Y - osb5}}, Line{&Point{c2.X - sb5, c2.Y - sa5}, &Point{c.X + osb7, c.Y + osa7}}
}

// Note: more than 180 degree not always possible and rotation is left to right +-180
// and two turn, stopAngle might be beyond a full turn...
func (c *Circle) Coordinates() []Point {
    startAngle := rToD(math.Atan2(c.StartPoint.Y-c.Y, c.StartPoint.X-c.X))
    stopAngle := rToD(math.Atan2(c.StopPoint.Y-c.Y, c.StopPoint.X-c.X))

    points := []Point{}
    print := false
out:
    for i := 0; i < 2; i++ { 
        for i := 0; i <= 180; i++ {
            if !print && startAngle >= 0.0 && startAngle <= float64(i) {
                points = append(points, c.PointInAngle(startAngle))
                print = true
            }
            if print && stopAngle >= 0.0 && stopAngle <= float64(i) {
                points = append(points, c.PointInAngle(stopAngle))
                break out
            }
            if print {
                points = append(points, c.PointInAngle(float64(i)))
            }
        }
        for i := -179; i < 0; i++ {
            if !print && startAngle < 0.0 && startAngle <= float64(i) {
                points = append(points, c.PointInAngle(startAngle))
                print = true
            }
            if print && stopAngle < 0.0 && stopAngle <= float64(i) {
                points = append(points, c.PointInAngle(stopAngle))
                break out
            }
            if print {
                points = append(points, c.PointInAngle(float64(i)))
            }
        }
    }

    return points
}

func newCircle(x, y, r float64) *Circle {
    return &Circle{&Point{x, y}, r, &Point{x, y}, &Point{x, y}}
}

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

