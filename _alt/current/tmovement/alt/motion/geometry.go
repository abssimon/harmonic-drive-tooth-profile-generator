package main

import (
    "math"
)

// Point

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

// Line

type Line struct {
    p1 *Point
    p2 *Point
}

// Cirle



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
    sa3 := math.Sin(rad1) * c2.Radius
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
        for i := -179; i <= 0; i++ {
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

func (e Ellipse) PointInAngle(a float64) Point {
    da := dToR(a)
    x := e.Width * math.Cos(da)
    y := e.Height * math.Sin(da)
    return Point{x, y}
}

// https://math.stackexchange.com/questions/2645689
func (e Ellipse) PointByAngleRotated(a float64, r float64) Point {
    da := dToR(a)
    dr := dToR(r)
    cda := math.Cos(da)
    cdr := math.Cos(dr)
    sda := math.Sin(da)
    sdr := math.Sin(dr)
    x := e.Width*cda*cdr - e.Height*sda*sdr
    y := e.Width*cda*sdr + e.Height*sda*cdr
    return Point{x, y}
}

// https://math.stackexchange.com/questions/4086824
func (e Ellipse) Tangent(a float64) float64 {
    return math.Atan2(e.Height*math.Cos(dToR(a)), -e.Width*math.Sin(dToR(a)))
}

// TangentByPoint is used to calculate a tangens for a rotated ellipse.
// A point from a rotated ellipse can be rotated back, and with this
// a normal tangens can be calculated
// https://math.stackexchange.com/questions/4134884
func (e Ellipse) TangentByPoint(p Point, i float64) float64 {
    a := math.Atan2(p.Y/e.Height, p.X/e.Width)
    return math.Atan2(e.Height*math.Cos(a), -e.Width*math.Sin(a))
}
