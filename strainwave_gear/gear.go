package main

import (
    "math"
)

type Gear struct {
    Tooths []Tooth
    Angle float64
}

type Tooth struct {
    C1 *Circle
    C2 *Circle
    C3 *Circle
    C4 *Circle
    pos Point
}

func (t Tooth) Rotate(a float64) {
    t.C1.Rotate(a)
    t.C2.Rotate(a)
    t.C3.Rotate(a)
    t.C4.Rotate(a)
}

// A s-tooth is definfined by segments of circles. Start and stop are angles, used
// to define these segments. Only an inner tangent is needed to connect these segments 
// to a path, because the tooth is symmetrical
func NewTooth(conf *Config, rotate float64, bottomStop float64, p Point) Tooth {

    tipCenter := Point{conf.TipCenterX * conf.Scale, conf.TipCenterY * conf.Scale}
    tipRadius := conf.TipRadius * conf.Scale
    tipStop := conf.TipStop 
    bottomCenter := Point{conf.BottomCenterX * conf.Scale, conf.BottomCenterY * conf.Scale}
    bottomRadius := conf.BottomRadius * conf.Scale
    
    tooth := Tooth{
        &Circle{&Point{-bottomCenter.X, bottomCenter.Y}, bottomRadius, 0.0, 0.0},
        &Circle{&Point{tipCenter.X, tipCenter.Y}, tipRadius, 0.0, 0.0},
        &Circle{&Point{-tipCenter.X, tipCenter.Y}, tipRadius, 0.0, 0.0},
        &Circle{&Point{bottomCenter.X, bottomCenter.Y}, bottomRadius, 0.0, 0.0},
        p,
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

    tooth.Rotate(rotate)

    return tooth
}

