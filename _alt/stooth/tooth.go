package main

import (
    "math"
)

type Tooth struct {
    C1   *Circle
    C2   *Circle
    C3   *Circle
    C4   *Circle
}

func (t Tooth) rotate(a float64) {
    t.C1.Rotate(a)
    t.C2.Rotate(a)
    t.C3.Rotate(a)
    t.C4.Rotate(a)
}

func EllipseTooth(e Ellipse, i float64, p Point) []Point {
    
    // tooth definition
    scale := 6.0
    tipCenter := Point{-0.06 * scale, 1.453 * scale}
    tipRadius := 0.356 * scale
    tipStop := 1.56 // 90 grad
    bottomCenter := Point{1.633 * scale, 1.176 * scale}
    bottomRadius := 0.85 * scale
    bottomStop := 1.56 // 90 grad
    
    tooth := Tooth{
        &Circle{&Point{-bottomCenter.X, bottomCenter.Y}, bottomRadius, 0.0, 0.0},
        &Circle{&Point{tipCenter.X, tipCenter.Y}, tipRadius, 0.0, 0.0},
        &Circle{&Point{-tipCenter.X, tipCenter.Y}, tipRadius, 0.0, 0.0},
        &Circle{&Point{bottomCenter.X, bottomCenter.Y}, bottomRadius, 0.0, 0.0},
    }

    // symmetrical, just need one tan
    _, tan := tooth.C3.InnerTangentWith(tooth.C4)
    tooth.C1.Start = math.Pi*2.0 - bottomStop
    tooth.C1.Stop = math.Pi*2.0 - tan
    tooth.C2.Start = math.Pi - tipStop
    tooth.C2.Stop = math.Pi - tan
    tooth.C3.Start = tan
    tooth.C3.Stop = tipStop
    tooth.C4.Start = math.Pi + tan
    tooth.C4.Stop = math.Pi + bottomStop
    
    /*
    tipCenter := Point{2.333 * scale, 1.176 * scale}
    tipRadius := 0.866 * scale
    tipStop := 0.925 // 53
    bottomCenter := Point{0.24 * scale, 1.453 * scale}
    bottomRadius := 1.67 * scale
    bottomStop := 0.78 // 45

    tooth := Tooth{
        &Circle{&Point{-tipCenter.X, tipCenter.Y}, bottomRadius, 0.0, 0.0},
        &Circle{&Point{bottomCenter.X, bottomCenter.Y}, tipRadius, 0.0, 0.0},
        &Circle{&Point{-bottomCenter.X, bottomCenter.Y}, tipRadius, 0.0, 0.0},
        &Circle{&Point{tipCenter.X, tipCenter.Y}, bottomRadius, 0.0, 0.0},
    }

    // symmetrical, just need one tan
    _, tan := tooth.C3.InnerTangentWith(tooth.C4)
    tooth.C1.Start = math.Pi*2.0 - bottomStop
    tooth.C1.Stop = math.Pi*2.0 - tan
    tooth.C2.Start = math.Pi - tipStop
    tooth.C2.Stop = math.Pi - tan
    tooth.C3.Start = tan
    tooth.C3.Stop = tipStop
    tooth.C4.Start = math.Pi + tan
    tooth.C4.Stop = math.Pi + bottomStop
    */
    // rotate correctly for ellipse
    tan = e.Tangent(i)  // is i correct?
    tooth.rotate(math.Pi + tan)
    
    // append points in order
    points := []Point{}
    c := tooth.C4.Coordinates()
    for i := len(c) - 1; i >= 0; i-- {
        points = append(points, Point{c[i].X+p.X, c[i].Y+p.Y})        
    }
    c = tooth.C3.Coordinates()
    for _, point := range c {
        points = append(points, Point{point.X+p.X, point.Y+p.Y}) 
    }
    c = tooth.C2.Coordinates()
    for _, point := range c {
        points = append(points, Point{point.X+p.X, point.Y+p.Y})     
    }
    c = tooth.C1.Coordinates()
    for i := len(c) - 1; i >= 0; i-- {
        points = append(points, Point{c[i].X+p.X, c[i].Y+p.Y}) 
    }
    
    return points
}