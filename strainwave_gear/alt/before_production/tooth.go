package main

import (
    "math"
)

type Tooth struct {
    C1 *Circle
    C2 *Circle
    C3 *Circle
    C4 *Circle
}

func (t Tooth) rotate(a float64) {
    t.C1.Rotate(a)
    t.C2.Rotate(a)
    t.C3.Rotate(a)
    t.C4.Rotate(a)
}

// A s-tooth is definfined by parts of circles, start and stop are the angles
// to define the parts. Only one inner tangent is needed, because the tooth is 
// symmetrical
func BasicTooth(rotate float64) Tooth {

    // tooth definition
    scale := 7.0
    tipCenter := Point{-0.06 * scale, 1.453 * scale}
    tipRadius := 0.356 * scale
    tipStop := 1.56 // 90 gr
    bottomCenter := Point{1.633 * scale, 1.176 * scale}
    bottomRadius := 0.85 * scale
    bottomStop := 1.56 // 90 grad

    tooth := Tooth{
        &Circle{&Point{-bottomCenter.X, bottomCenter.Y}, bottomRadius, 0.0, 0.0},
        &Circle{&Point{tipCenter.X, tipCenter.Y}, tipRadius, 0.0, 0.0},
        &Circle{&Point{-tipCenter.X, tipCenter.Y}, tipRadius, 0.0, 0.0},
        &Circle{&Point{bottomCenter.X, bottomCenter.Y}, bottomRadius, 0.0, 0.0},
    }

    _, tan := tooth.C3.InnerTangentWith(tooth.C4)
    tooth.C1.Start = math.Pi*2.0 - bottomStop
    tooth.C1.Stop = math.Pi*2.0 - tan
    tooth.C2.Start = math.Pi - tipStop
    tooth.C2.Stop = math.Pi - tan
    tooth.C3.Start = tan
    tooth.C3.Stop = tipStop
    tooth.C4.Start = math.Pi + tan
    tooth.C4.Stop = math.Pi + bottomStop

    tooth.rotate(rotate)

    return tooth
}

// 
func FlexTooth(p Point, rotate float64) []Point {

    tooth := BasicTooth(rotate)

    // append points in order
    points := []Point{}
    c := tooth.C4.Coordinates()
    for i := len(c) - 1; i >= 0; i-- {
        points = append(points, Point{c[i].X + p.X, c[i].Y + p.Y})
    }
    c = tooth.C3.Coordinates()
    for _, point := range c {
        points = append(points, Point{point.X + p.X, point.Y + p.Y})
    }
    c = tooth.C2.Coordinates()
    for _, point := range c {
        points = append(points, Point{point.X + p.X, point.Y + p.Y})
    }
    c = tooth.C1.Coordinates()
    for i := len(c) - 1; i >= 0; i-- {
        points = append(points, Point{c[i].X + p.X, c[i].Y + p.Y})
    }

    return points
}

// Ridid Tooth is a copy of a Flextooth, but without rounded
// parts. But flank angles are the same.
func RigidTooth(p Point, rotate float64) []Point {

    tooth := BasicTooth(rotate)

    // append points in order, skip round bottom part
    points := []Point{}
    p1 := tooth.C4.PointInAngle(tooth.C4.Start)
    p2 := tooth.C3.PointInAngle(tooth.C3.Start)
    p3 := tooth.C2.PointInAngle(tooth.C2.Stop)
    p4 := tooth.C1.PointInAngle(tooth.C1.Stop)

    // make flank longer and calculate intersection 
    // instead of a round top
    p2.X += (p2.X - p1.X) * 10.0
    p2.Y += (p2.Y - p1.Y) * 10.0
    p3.X += (p3.X - p4.X) * 10.0
    p3.Y += (p3.Y - p4.Y) * 10.0
    ip, _ := intersection(p1, p2, p3, p4)

    points = append(points, Point{p1.X + p.X, p1.Y + p.Y})
    points = append(points, Point{ip.X + p.X, ip.Y + p.Y})
    points = append(points, Point{p4.X + p.X, p4.Y + p.Y})

    return points
}
