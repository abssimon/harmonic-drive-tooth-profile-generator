package main

import (
	"errors"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) DistanceTo(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Line struct {
	p1 Point
	p2 Point
}

type Circle struct {
	Point  // kann man auch als ganzes ansprechen, oder Center
	Radius float64
}

func (c Circle) PointInAngle(a float64) Point {
	x := math.Cos(dToR(a)) * c.Radius
	y := math.Sin(dToR(a)) * c.Radius

	return Point{x + c.X, y + c.Y}
}

func (c *Circle) OuterTangentTo(c2 *Circle) (Line, Line) {
	distCenterToCenter := c.Point.DistanceTo(c2.Point) // math.Sqrt(math.Pow(c.Y-c2.Y, 2) + math.Pow(c.2.X-cX, 2))

	// Outer
	radiusDiff := math.Abs(c2.Radius - c.Radius)
	diffRatio := distCenterToCenter / radiusDiff
	distTanIntersection := diffRatio * c.Radius

	alpha1 := rToD(math.Asin(c.Radius / distTanIntersection))
	alpha2 := rToD(math.Asin((c.Y-c2.Y)/distCenterToCenter)) - alpha1
	alpha3 := math.Asin((c.Y - c2.Y) / distCenterToCenter)
	alpha4 := 90 - (rToD(math.Asin((c.Y-c2.Y)/distCenterToCenter)) + alpha1)

	c1 := math.Sqrt(math.Pow(distTanIntersection, 2) - math.Pow(c.Radius, 2))
	a1 := math.Sin(dToR(alpha2)) * c1
	b1 := math.Cos(dToR(alpha2)) * c1

	c2_ := distCenterToCenter + distTanIntersection
	b2 := math.Cos(alpha3) * c2_
	a2 := math.Sin(alpha3) * c2_

	b3 := math.Sqrt(math.Pow(c2_, 2) - math.Pow(c2.Radius, 2))
	b4 := math.Cos(dToR(alpha2)) * b3
	a4 := math.Sin(dToR(alpha2)) * b3
	b5 := math.Cos(dToR(alpha4)) * b3
	a5 := math.Sin(dToR(alpha4)) * b3
	b6 := math.Cos(dToR(alpha4)) * c1
	a6 := math.Sin(dToR(alpha4)) * c1

	ttx1 := c2.X - b2 + b1
	tty1 := c2.Y + a2 - a1
	ttx2 := c2.X - b2 + b4
	tty2 := c2.Y + a2 - a4
	tbx1 := c2.X - b2 + a5
	tby1 := c2.Y + a2 - b5
	tbx2 := c2.X - b2 + a6
	tby2 := c2.Y + a2 - b6

	return Line{Point{tbx1, tby1}, Point{tbx2, tby2}}, Line{Point{ttx1, tty1}, Point{ttx2, tty2}}
}

// Tangent Inner Outer automatisch, wenn abstand der mittelpunkt und 2x radius berechnet wird
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

	sa3 := math.Sin(dToR(a3)) * c2.Radius // todo names
	sb3 := math.Cos(dToR(a3)) * c2.Radius
	sa5 := math.Sin(dToR(a4)) * c2.Radius
	sb5 := math.Cos(dToR(a4)) * c2.Radius

	osa5 := math.Sin(dToR(a3)) * c.Radius
	osb5 := math.Cos(dToR(a3)) * c.Radius
	osa7 := math.Sin(dToR(a4)) * c.Radius
	osb7 := math.Cos(dToR(a4)) * c.Radius

	return Line{Point{c2.X - sa3, c2.Y + sb3}, Point{c.X + osa5, c.Y - osb5}}, Line{Point{c2.X - sb5, c2.Y - sa5}, Point{c.X + osb7, c.Y + osa7}}
}

// Note: more than 180 degree not always possible and rotation is left to right +-180
func (c *Circle) CoordinatesBetween(start Point, stop Point) []Point {
	startAngle := rToD(math.Atan2(start.Y-c.Y, start.X-c.X))
	stopAngle := rToD(math.Atan2(stop.Y-c.Y, stop.X-c.X))

	points := []Point{}
	print := false
out:
	for i := 0; i < 2; i++ { // two turn, stopAngle might be beyond a full turn...
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

// https://github.com/rahuldhole/Trilateration/blob/master/Trilateration.php
func (c *Circle) InterectionPoints(c2 *Circle) (*Point, *Point, error) {
	deltaX := c.X - c2.X
	deltaY := c.Y - c2.Y
	delta := math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))

	if delta > (c.Radius + c2.Radius) {
		return nil, nil, errors.New("There is no intersection")
	} else if delta < (c.Radius - c2.Radius) {
		return nil, nil, errors.New("One circle is inside other")
	}

	s := (math.Pow(delta, 2) + math.Pow(c2.Radius, 2) - math.Pow(c.Radius, 2)) / (2.0 * delta)
	Rx := c2.X + (deltaX*s)/delta
	Ry := c2.Y + (deltaY*s)/delta
	u := math.Sqrt(math.Pow(c2.Radius, 2) - math.Pow(s, 2))

	return &Point{Rx - (deltaY*u)/delta, Ry + (deltaX*u)/delta}, &Point{Rx + (deltaY*u)/delta, Ry - (deltaX*u)/delta}, nil
}

// Create a circular object
func newCircle(x, y, r float64) *Circle {
	return &Circle{Point{x, y}, r}
}
