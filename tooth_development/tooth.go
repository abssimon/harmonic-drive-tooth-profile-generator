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

func EllipseTooth(e Ellipse, i float64, p Point, scale float64) Tooth {
    
    // tooth definition
    tipCenter := Point{-0.0 * scale, 1.265 * scale}
    tipRadius := 0.506 * scale
    tipStop := 1.56 // 90 grad
    bottomCenter := Point{1.305 * scale, 0.985 * scale}
    bottomRadius := 0.506 * scale
    bottomStop := 0.95 // 90 grad
    
    
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
    
    // rotate correctly for ellipse
    tan = e.Tangent(i)  // is i correct?
    tooth.rotate(math.Pi + tan)
    
    return tooth
}