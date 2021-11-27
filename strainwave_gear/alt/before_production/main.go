package main

import (
    "math"
)

// Teeth mashing is pretty hard. The only working solution is found here
// https://commons.wikimedia.org/wiki/File:HarmonicDriveAni.gif
func toothPosition(ellipseCoordinates []Point, nPoints float64, offset float64) []Point {

    // changes (tiny) of coordinates
    sum := make([]float64, len(ellipseCoordinates))
    for i, v := range Diff(ellipseCoordinates) {
        sum[i+1] = math.Hypot(v.X, v.Y)
    }
    cumLen := CumSum(sum)
    circumference := cumLen[len(cumLen)-1]
    //fmt.Println("entspricht kreis durchmesser ", circumference / math.Pi)

    // location on curcumference, with offset
    pos := LinSpace(0.0, 1.0, int(nPoints)+1)
    pos = pos[:len(pos)-1]
    for i, v := range pos {
        pos[i] = math.Mod(v+offset, 1) * circumference
    }

    coordX := make([]float64, len(ellipseCoordinates))
    coordY := make([]float64, len(ellipseCoordinates))
    for i, v := range ellipseCoordinates {
        coordX[i] = v.X
        coordY[i] = v.Y
    }

    // predict X for pos
    x, err := interp1(cumLen, coordX, pos)
    if err != nil {
        panic(err)
    }
    y, err := interp1(cumLen, coordY, pos)
    if err != nil {
        panic(err)
    }

    r := make([]Point, len(x))
    for i, v := range x {
        r[i] = Point{v, y[i]}
    }

    return r
}

func main() {

    rigidTheets := 102.0
    flexTheets := rigidTheets - 2.0

    nFrames := 1 // 100
    frameAngles := LinSpace(0.0, -math.Pi, nFrames+1)
    frameAngles = frameAngles[:len(frameAngles)-1]

    diameterH := 4.2
    diameterV := 4.035 // calculate?

    // sample points for prediction
    e := Ellipse{Point{0, 0}, diameterH, diameterV}
    ellipseCoordinates := e.Coordinates(1000)

    gears := [][]Point{}
    for _, angleWaveGen := range frameAngles {

        // position on ellipse
        angleFlexTeeth := angleWaveGen * (flexTheets - rigidTheets) / flexTheets // * -0,02
        offset := (-angleWaveGen + angleFlexTeeth) / 2.0 / math.Pi               
        r := toothPosition(ellipseCoordinates, flexTheets, offset)

        // add teeth
        gear := []Point{}
        for _, p := range r {
            sp := Point{p.X * 200.0, p.Y * 200.0}
            gear = append(gear, FlexTooth(sp, math.Pi+e.TangentByPoint(sp))...)
        }

        for i := range gear {
            gear[i].Rotate(angleWaveGen) 
        }

        gears = append(gears, gear)
    }
/*
    // draw rigid gear
    c := &Circle{&Point{0.0, 0.0}, diameterH / 2, 0.0, math.Pi * 2.0}
    angle := LinSpace(0.0, math.Pi*2.0, int(rigidTheets)+1)
    angle = angle[:len(angle)-1]
    gear := []Point{}
    for _, a := range angle {
        cp := c.PointInAngle(a)
        gear = append(gear, RigidTooth(Point{cp.X * 200.0, cp.Y * 200.0}, -math.Pi/2.0+a)...)
    }
    gears = append(gears, gear)
*/    
    // draw round ellipse for production
    re := &Circle{&Point{0.0, 0.0}, 4.117906474475928 / 2, 0.0, math.Pi * 2.0} // diameter calculation see above
    angle := LinSpace(0.0, math.Pi*2.0, int(flexTheets))
    angle = angle[:len(angle)-1]
    gear := []Point{}
    for _, a := range angle {
        cp := re.PointInAngle(a)
        gear = append(gear, RigidTooth(Point{cp.X * 200.0, cp.Y * 200.0}, -math.Pi/2.0+a)...) // TODO hier FlexTooth !!!
    }
    gears = append(gears, gear)
    
    /*
    // rotate all gears manually to have one flanc horizontal
    for i := range gears {
        for j := range gears[i] {
            gears[i][j].Rotate(0.05) 
        }
    }
    */
    
    // wie bekomme ich die ellipse rund?
    // kreis mit gleichem umfang finden?

    svg(gears)
}
